package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	planstypes "github.com/lavanet/lava/x/plans/types"
)

// AccountKeeper defines the expected account keeper used for simulations (noalias)
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) types.AccountI
	// Methods imported from account should be defined here
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error
	// Methods imported from bank should be defined here
}

type EpochstorageKeeper interface {
	GetEpochStart(ctx sdk.Context) uint64
	// Methods imported from epochstorage should be defined here
}

type ProjectsKeeper interface {
	CreateDefaultProject(ctx sdk.Context, consumer string) error
	DeleteProject(ctx sdk.Context, index string) error
	// Methods imported from projectskeeper should be defined here
}

type PlansKeeper interface {
	GetPlan(ctx sdk.Context, index string) (planstypes.Plan, bool)
	PutPlan(ctx sdk.Context, index string, block uint64)
	// Methods imported from planskeeper should be defined here
}
