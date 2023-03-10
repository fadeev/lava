package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/lavanet/lava/x/subscription/types"
)

// EpochStart runs the functions that are supposed to run in epoch start
func (k Keeper) EpochStart(ctx sdk.Context) {
	// On epoch start we need to iterate through all the subscriptions and
	// check those whose current month just expired:
	//
	// - record the actual epoch of expiry (for rewards validation)
	// - save the month's remaining CU in prev (for rewards validation)
	// - reset the month's remaining CU to the plan's allowance
	// - reduce remaining duration, and delete if it reaches zero
	//
	// Note that actual deletion is deferred by PAYMENT_GRACE_TIME_HOURS,
	// to allow for payments for the last month of the subscription.
	// (During this period, remaining CU of the new month is zero).

	date := ctx.BlockTime().UTC()
	subExpired := k.GetCondSubscription(ctx, func(sub types.Subscription) bool {
		return sub.IsMonthExpired(date)
	})

	// First collect, then remove (because removing may affect the iterator?)
	for _, sub := range subExpired {
		sub.PrevExpiryBlock = uint64(ctx.BlockHeight())
		sub.PrevCuLeft = sub.MonthCuLeft

		// subscription has been dead for a month: delete it
		if sub.DurationLeft == 0 {
			k.RemoveSubscription(ctx, sub.Consumer)
			continue
		}

		sub.DurationLeft -= 1

		if sub.DurationLeft > 0 {
			date = nextMonth(date)
			sub.MonthExpiryTime = uint64(date.Unix())
			sub.MonthCuLeft = sub.MonthCuTotal

			// reset CU allowance for this coming month
			sub.MonthCuLeft = sub.MonthCuTotal
		} else {
			// duration ended, but don't delete yet - keep around for another
			// PAYMENT_GRACE_TIME_HOURS blocks before removing, to allow for
			// payments for the months the ends now to be validated.
			date = date.Add(time.Hour * types.PAYMENT_GRACE_TIME_HOURS)
			sub.MonthExpiryTime = uint64(date.Unix())

			// zero CU allowance for this coming month
			sub.MonthCuLeft = 0
		}

		k.SetSubscription(ctx, sub)
	}
}
