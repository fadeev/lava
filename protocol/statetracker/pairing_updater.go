package statetracker

import (
	"fmt"
	"strconv"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/protocol/lavasession"
	"github.com/lavanet/lava/utils"
	epochstoragetypes "github.com/lavanet/lava/x/epochstorage/types"
	"golang.org/x/net/context"
)

const (
	CallbackKeyForPairingUpdate = "pairing-update"
)

type PairingUpdater struct {
	consumerSessionManagersMap map[string][]*lavasession.ConsumerSessionManager // key is chainID so we don;t run getPairing more than once per chain
	nextBlockForUpdate         uint64
	stateQuery                 *ConsumerStateQuery
}

func NewPairingUpdater(consumerAddress sdk.AccAddress, stateQuery *ConsumerStateQuery) *PairingUpdater {
	return &PairingUpdater{consumerSessionManagersMap: map[string][]*lavasession.ConsumerSessionManager{}, stateQuery: stateQuery}
}

func (pu *PairingUpdater) RegisterPairing(ctx context.Context, consumerSessionManager *lavasession.ConsumerSessionManager) error {
	chainID := consumerSessionManager.RPCEndpoint().ChainID
	pairingList, epoch, nextBlockForUpdate, err := pu.stateQuery.GetPairing(context.Background(), chainID, -1)
	if err != nil {
		return err
	}
	pu.updateConsummerSessionManager(ctx, pairingList, consumerSessionManager, epoch)
	pu.nextBlockForUpdate = nextBlockForUpdate // make sure we don't update twice when launching.
	consumerSessionsManagersList, ok := pu.consumerSessionManagersMap[chainID]
	if !ok {
		pu.consumerSessionManagersMap[chainID] = []*lavasession.ConsumerSessionManager{consumerSessionManager}
		return nil
	}
	pu.consumerSessionManagersMap[chainID] = append(consumerSessionsManagersList, consumerSessionManager)
	return nil
}

func (pu *PairingUpdater) UpdaterKey() string {
	return CallbackKeyForPairingUpdate
}

func (pu *PairingUpdater) Update(latestBlock int64) {
	ctx := context.Background()
	if int64(pu.nextBlockForUpdate) > latestBlock {
		return
	}
	nextBlockForUpdateList := []uint64{}
	for chainID, consumerSessionManagerList := range pu.consumerSessionManagersMap {
		pairingList, epoch, nextBlockForUpdate, err := pu.stateQuery.GetPairing(ctx, chainID, latestBlock)
		if err != nil {
			utils.LavaFormatError("could not update pairing for chain, trying again next block", err, &map[string]string{"chain": chainID})
			nextBlockForUpdateList = append(nextBlockForUpdateList, pu.nextBlockForUpdate+1)
			continue
		} else {
			nextBlockForUpdateList = append(nextBlockForUpdateList, nextBlockForUpdate)
		}
		for _, consumerSessionManager := range consumerSessionManagerList {
			// same pairing for all apiInterfaces, they pick the right endpoints from inside using our filter function
			err := pu.updateConsummerSessionManager(ctx, pairingList, consumerSessionManager, epoch)
			if err != nil {
				utils.LavaFormatError("failed updating consumer session manager", err, &map[string]string{"chainID": chainID, "apiInterface": consumerSessionManager.RPCEndpoint().ApiInterface, "pairingListLen": strconv.Itoa(len(pairingList))})
				continue
			}
		}
	}
	nextBlockForUpdateMin := uint64(0)
	for idx, blockToUpdate := range nextBlockForUpdateList {
		if idx == 0 || blockToUpdate < nextBlockForUpdateMin {
			nextBlockForUpdateMin = blockToUpdate
		}
	}
	pu.nextBlockForUpdate = nextBlockForUpdateMin
}

func (pu *PairingUpdater) updateConsummerSessionManager(ctx context.Context, pairingList []epochstoragetypes.StakeEntry, consumerSessionManager *lavasession.ConsumerSessionManager, epoch uint64) (err error) {
	pairingListForThisCSM, err := pu.filterPairingListByEndpoint(ctx, pairingList, consumerSessionManager.RPCEndpoint(), epoch)
	if err != nil {
		return err
	}
	err = consumerSessionManager.UpdateAllProviders(epoch, pairingListForThisCSM)
	return
}

func (pu *PairingUpdater) filterPairingListByEndpoint(ctx context.Context, pairingList []epochstoragetypes.StakeEntry, rpcEndpoint lavasession.RPCEndpoint, epoch uint64) (filteredList []*lavasession.ConsumerSessionsWithProvider, err error) {
	// go over stake entries, and filter endpoints that match geolocation and api interface
	pairing := []*lavasession.ConsumerSessionsWithProvider{}
	for _, provider := range pairingList {
		//
		// Sanity
		providerEndpoints := provider.GetEndpoints()
		if len(providerEndpoints) == 0 {
			utils.LavaFormatError("skipping provider with no endoints", nil, &map[string]string{"Address": provider.Address, "ChainID": provider.Chain})
			continue
		}

		relevantEndpoints := []epochstoragetypes.Endpoint{}
		for _, endpoint := range providerEndpoints {
			// only take into account endpoints that use the same api interface and the same geolocation
			if endpoint.UseType == rpcEndpoint.ApiInterface && endpoint.Geolocation == rpcEndpoint.Geolocation {
				relevantEndpoints = append(relevantEndpoints, endpoint)
			}
		}
		if len(relevantEndpoints) == 0 {
			utils.LavaFormatError("skipping provider, No relevant endpoints for apiInterface", nil, &map[string]string{"Address": provider.Address, "ChainID": provider.Chain, "apiInterface": rpcEndpoint.ApiInterface, "Endpoints": fmt.Sprintf("%v", providerEndpoints)})
			continue
		}

		maxcu, err := pu.stateQuery.GetMaxCUForUser(ctx, provider.Chain, epoch)
		if err != nil {
			return nil, err
		}
		//
		pairingEndpoints := make([]*lavasession.Endpoint, len(relevantEndpoints))
		for idx, relevantEndpoint := range relevantEndpoints {
			endp := &lavasession.Endpoint{NetworkAddress: relevantEndpoint.IPPORT, Enabled: true, Client: nil, ConnectionRefusals: 0}
			pairingEndpoints[idx] = endp
		}

		pairing = append(pairing, &lavasession.ConsumerSessionsWithProvider{
			PublicLavaAddress: provider.Address,
			Endpoints:         pairingEndpoints,
			Sessions:          map[int64]*lavasession.SingleConsumerSession{},
			MaxComputeUnits:   maxcu,
			ReliabilitySent:   false,
			PairingEpoch:      epoch,
		})
	}
	if len(pairing) == 0 {
		return nil, utils.LavaFormatError("Failed getting pairing for consumer, pairing is empty", err, &map[string]string{"apiInterface": rpcEndpoint.ApiInterface, "ChainID": rpcEndpoint.ChainID, "geolocation": strconv.FormatUint(rpcEndpoint.Geolocation, 10)})
	}
	// replace previous pairing with new providers
	return pairing, nil
}
