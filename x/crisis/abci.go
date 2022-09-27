package crisis

import (
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
)

// check all registered invariants
func EndBlocker(ctx sdk.Context, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	haltWhenNecessary(ctx, k)

	if k.InvCheckPeriod() == 0 || ctx.BlockHeight()%int64(k.InvCheckPeriod()) != 0 {
		// skip running the invariant check
		return
	}
	k.AssertInvariants(ctx)

	haltWhenNecessary(ctx, k)
}

// haltWhenNecssary will halt the chain if it is configured to do so AND an invariant has failed
func haltWhenNecessary(ctx sdk.Context, k keeper.Keeper) {
	if k.InvHaltNode() && k.GetMustHalt(ctx) { // An invariant is broken AND the chain is configured to halt on an invariant failure
		panic("Crisis module: invariant broken - chain is halting immediately!")
	}
}
