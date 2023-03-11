package common

import (
	"math"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/common/types"
	"github.com/lavanet/lava/utils"
)

// FixationStore manages lists of entries with versions in the store.
// (See also documentation in common/fixation_entry_index.go)
//
// Its primary use it to implemented "fixated entries": entries that may change over
// time, and whose versions must be retained on-chain as long as they are referenced.
// For examples, an older version of a package is needed as long as the subscription
// that uses it lives.
//
// Once instantiated with NewFixationStore(), it offers the following methods:
//    - AppendEntry(index, block, *entry): add a new "block" version of an entry "index".
//    - ModifyEntry(index, block, *entry): modify existing entry with "index" and "block"
//    - FindEntry(index, block, *entry): get a copy (no reference) of a version of an entry
//    - GetEntry(index, *entry): get a copy (and reference) of the latest version of an entry
//    - PutEntry(index, block): drop a reference of a version of an entry
//    - [TBD] RemoveEntry(index): mark an entry as unavailable for new GetEntry() calls
//    - GetAllEntryIndex(): get all the entries indices (without versions)
//
// Entry names (index) must contain only visible ascii characters (ascii values 32-125).
// The ascii 'DEL' invisible character is used internally to terminate the index values
// when stored, to ensure that no two indices can ever overlap, i.e. one being the prefix
// of the other.
// (Note that this properly supports Bech32 addresses, which are limited to use only
// visible ascii characters as per https://en.bitcoin.it/wiki/BIP_0173#Specification).
//
// How does it work? The explanation below illustrates how the data is stored, assuming the
// user is the module "packages":
//
// 1. When instantiated, FixationStore gets a `prefix string` - used as a namespace
// to separate between instances of FixationStore. For instance, module "packages"
// would use its module name for prefix.
//
// 2. Each entry is wrapped by a RawEntry, that holds the index, block, ref-count,
// and the (marshalled) data of the entry.
//
// 3. FixationStore keeps the entry indices with a special prefix; and it keeps the
// and the RawEntres with a prefix that includes their index (using the block as key).
// For instance, modules "packages" may have a package named "YourPackage" created at
// block 110, and it was updated at block 220, and package named "MyPakcage" created
// at block 150. The store with prefix "packages" will hold the following objects:
//
//     prefix: packages_Entry_Index_            key: MyPackage      data: MyPackage
//     prefix: packages_Entry_Index_            key: YourPackage    data: YourPackage
//     prefix: packages_Entry_Value_MyPackage     key: 150            data: Entry
//     prefix: packages_Entry_Value_YourPackage   key: 220            data: Entry
//     prefix: packages_Entry_Value_YourPackage   key: 110            data: Entry
//
// Thus, iterating on the prefix "packages_Entry_Index_" would yield all the package
// indices. Reverse iterating on the prefix "packages_Entry_Raw_<INDEX>" would yield
// all the Fixation of the entry named <INDEX> in descending order.
//
// 4. FixationStore keeps a reference count of Fixation of entries, and when the
// count reaches 0 it marks them for deletion. The actual deletion takes place after
// a fixed number of epochs has passed.

type FixationStore struct {
	storeKey sdk.StoreKey
	cdc      codec.BinaryCodec
	prefix   string
}

// sanitizeIdnex checks that a string contains only visible ascii characters
// (i.e. Ascii 32-126), and appends a (ascii) DEL to the index; this ensures
// that an index can never be a prefix of another index.
func sanitizeIndex(index string) (string, error) {
	const (
		ASCII_MIN = 32  // min visible ascii
		ASCII_MAX = 126 // max visible ascii
		ASCII_DEL = 127 // ascii for DEL
	)

	for i := 0; i < len(index); i++ {
		if index[i] < ASCII_MIN || index[i] > ASCII_MAX {
			return index, types.ErrInvalidIndex
		}
	}
	return index + string([]byte{ASCII_DEL}), nil
}

// desantizeIndex reverts the effect of sanitizeIndex - removes the trailing
// (ascii) DEL terminator.
func desanitizeIndex(safeIndex string) string {
	return safeIndex[0 : len(safeIndex)-1]
}

func (fs *FixationStore) getStore(ctx sdk.Context, index string) *prefix.Store {
	store := prefix.NewStore(
		ctx.KVStore(fs.storeKey),
		types.KeyPrefix(fs.createStoreKey(index)))
	return &store
}

// setEntry modifies an existing entry in the store
func (fs *FixationStore) setEntry(ctx sdk.Context, entry types.Entry) {
	store := fs.getStore(ctx, entry.Index)
	byteKey := types.EncodeKey(entry.Block)
	marshaledEntry := fs.cdc.MustMarshal(&entry)
	store.Set(byteKey, marshaledEntry)
}

// AppendEntry adds a new entry to the store
func (fs *FixationStore) AppendEntry(ctx sdk.Context, index string, block uint64, entryData codec.ProtoMarshaler) error {
	safeIndex, err := sanitizeIndex(index)
	if err != nil {
		details := map[string]string{"index": index}
		return utils.LavaError(ctx, ctx.Logger(), "AppendEntry_invalid_index", details, "invalid non-ascii entry")
	}

	latestEntry, found := fs.getUnmarshaledEntryForBlock(ctx, safeIndex, block)

	// if latest entry is not found, this is a first version entry
	if !found {
		fs.setEntryIndex(ctx, safeIndex)
	} else {
		// make sure the new entry's block is not smaller than the latest entry's block
		if block < latestEntry.GetBlock() {
			details := map[string]string{
				"latestBlock": strconv.FormatUint(latestEntry.GetBlock(), 10),
				"block":       strconv.FormatUint(block, 10),
				"index":       index,
				"fs.prefix":   fs.prefix,
			}
			return utils.LavaError(ctx, ctx.Logger(), "AppendEntry_block_too_early", details, "entry block earlier than latest entry")
		}

		// if the new entry's block is equal to the latest entry, overwrite the latest entry
		if block == latestEntry.GetBlock() {
			return fs.ModifyEntry(ctx, index, block, entryData)
		}

		// if the old latest entry has refcount of 0, then update its "stale_at" time
		// TODO: remove this when the latest entry gets its own refcount.
		if latestEntry.Refcount == 0 {
			// never overflows because because ctx.BlockHeight is int64
			latestEntry.StaleAt = uint64(ctx.BlockHeight()) + uint64(types.STALE_ENTRY_TIME)
			fs.setEntry(ctx, latestEntry)
		}
	}

	// marshal the new entry's data
	marshaledEntryData := fs.cdc.MustMarshal(entryData)

	// create a new entry and marshal it
	entry := types.Entry{
		Index:    safeIndex,
		Block:    block,
		StaleAt:  math.MaxUint64,
		Data:     marshaledEntryData,
		Refcount: 0,
	}

	fs.setEntry(ctx, entry)

	fs.deleteStaleEntries(ctx, safeIndex)

	return nil
}

func (fs *FixationStore) deleteStaleEntries(ctx sdk.Context, safeIndex string) {
	// get the relevant store and init an iterator
	store := fs.getStore(ctx, safeIndex)
	iterator := sdk.KVStorePrefixIterator(store, []byte{})
	defer iterator.Close()

	// iterate over entries
	for iterator.Valid() {
		// umarshal the old entry version
		var oldEntry types.Entry
		fs.cdc.MustUnmarshal(iterator.Value(), &oldEntry)

		iterator.Next()

		// skipping removal of latest version
		if !iterator.Valid() {
			break
		}

		// delete stale entries (if they are at the end of the list)
		if oldEntry.IsStale(ctx) {
			fs.removeEntry(ctx, oldEntry.GetIndex(), oldEntry.GetBlock())
		} else {
			// avoid removal of entries in the middle of the list, because it would
			// break future lookup that may involve that (stale) entry.
			break
		}
	}
}

// ModifyEntry modifies an existing entry in the store
func (fs *FixationStore) ModifyEntry(ctx sdk.Context, index string, block uint64, entryData codec.ProtoMarshaler) error {
	safeIndex, err := sanitizeIndex(index)
	if err != nil {
		details := map[string]string{"index": index}
		return utils.LavaError(ctx, ctx.Logger(), "ModifyEntry_invalid_index", details, "invalid non-ascii entry")
	}

	// get the entry from the store
	entry, found := fs.getUnmarshaledEntryForBlock(ctx, safeIndex, block)
	if !found {
		details := map[string]string{
			"fs.prefix": fs.prefix,
			"index":     index,
			"block":     strconv.FormatUint(block, 10),
		}
		return utils.LavaError(ctx, ctx.Logger(), "SetEntry_cant_find_entry", details, "entry does not exist")
	}

	// update entry data
	marshaledEntryData := fs.cdc.MustMarshal(entryData)
	entry.Data = marshaledEntryData

	fs.setEntry(ctx, entry)
	return nil
}

// getUnmarshaledEntryForBlock gets an entry version for an index that has
// nearest-smaller block version for the given block arg.
func (fs *FixationStore) getUnmarshaledEntryForBlock(ctx sdk.Context, index string, block uint64) (types.Entry, bool) {
	store := fs.getStore(ctx, index)

	// init a reverse iterator
	iterator := sdk.KVStoreReversePrefixIterator(store, []byte{})
	defer iterator.Close()

	// iterate over entries in reverse order of version (block), and return the
	// first version that is smaller-or-equal to the requested block. Thus, the
	// caller gets the latest version at the time of the block (arg).
	// For example, assuming two entries for blocks 100 and 200, asking for the
	// entry with block 199 would return the entry with block 100, which was in
	// effect at the time of block 199.

	for ; iterator.Valid(); iterator.Next() {
		var entry types.Entry
		fs.cdc.MustUnmarshal(iterator.Value(), &entry)

		if entry.GetBlock() <= block {
			if entry.IsStale(ctx) {
				break
			}

			return entry, true
		}
	}
	return types.Entry{}, false
}

// FindEntry returns the entry by index and block without changing the refcount
func (fs *FixationStore) FindEntry(ctx sdk.Context, index string, block uint64, entryData codec.ProtoMarshaler) bool {
	safeIndex, err := sanitizeIndex(index)
	if err != nil {
		details := map[string]string{"index": index}
		utils.LavaError(ctx, ctx.Logger(), "FindEntry_invalid_index", details, "invalid non-ascii entry")
		return false
	}

	// get the unmarshaled entry for block
	entry, found := fs.getUnmarshaledEntryForBlock(ctx, safeIndex, block)
	if !found {
		return false
	}

	// unmarshal the entry's data
	fs.cdc.MustUnmarshal(entry.GetData(), entryData)

	return true
}

// GetEntry returns the latest entry by index and increments the refcount
func (fs *FixationStore) GetEntry(ctx sdk.Context, index string, entryData codec.ProtoMarshaler) bool {
	safeIndex, err := sanitizeIndex(index)
	if err != nil {
		details := map[string]string{"index": index}
		utils.LavaError(ctx, ctx.Logger(), "GetEntry_invalid_index", details, "invalid non-ascii entry")
		return false
	}

	block := uint64(ctx.BlockHeight())

	// get the unmarshaled entry for block
	entry, found := fs.getUnmarshaledEntryForBlock(ctx, safeIndex, block)
	if !found {
		return false
	}

	// unmarshal the entry's data
	fs.cdc.MustUnmarshal(entry.GetData(), entryData)

	entry.Refcount += 1
	fs.setEntry(ctx, entry)

	return true
}

// PutEntry finds the entry by index and block and decrements the refcount
func (fs *FixationStore) PutEntry(ctx sdk.Context, index string, block uint64) {
	safeIndex, err := sanitizeIndex(index)
	if err != nil {
		details := map[string]string{"index": index}
		utils.LavaError(ctx, ctx.Logger(), "PutEntry_invalid_index", details, "invalid non-ascii entry")
		return
	}

	// get the unmarshaled entry for block
	entry, found := fs.getUnmarshaledEntryForBlock(ctx, safeIndex, block)
	if !found {
		return
	}

	if entry.Refcount == 0 {
		// TODO: this indicates a bad bug ... panic?!
		details := map[string]string{
			"index":    index,
			"refcount": strconv.Itoa(int(entry.Refcount)),
		}
		utils.LavaError(ctx, ctx.Logger(), "PutEntry_decrement_count", details, "refcount already reached zero!")
		return
	}

	entry.Refcount -= 1

	if entry.Refcount == 0 {
		// never overflows because because ctx.BlockHeight is int64
		entry.StaleAt = uint64(ctx.BlockHeight()) + uint64(types.STALE_ENTRY_TIME) - 1
	}

	fs.setEntry(ctx, entry)
}

// removeEntry removes an entry from the store
func (fs *FixationStore) removeEntry(ctx sdk.Context, index string, block uint64) {
	store := fs.getStore(ctx, index)
	store.Delete(types.EncodeKey(block))
}

func (fs *FixationStore) createStoreKey(index string) string {
	return types.EntryKey + fs.prefix + index
}

// NewFixationStore returns a new FixationStore object
func NewFixationStore(storeKey sdk.StoreKey, cdc codec.BinaryCodec, prefix string) *FixationStore {
	fs := FixationStore{storeKey: storeKey, cdc: cdc, prefix: prefix}
	return &fs
}
