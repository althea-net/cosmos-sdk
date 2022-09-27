package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
)

// GetConstantFee gets the constant fee from the paramSpace
func (k Keeper) GetConstantFee(ctx sdk.Context) (constantFee sdk.Coin) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyConstantFee, &constantFee)
	return
}

// SetConstantFee sets the constant fee in the paramSpace
func (k Keeper) SetConstantFee(ctx sdk.Context, constantFee sdk.Coin) {
	k.paramSpace.Set(ctx, types.ParamStoreKeyConstantFee, constantFee)
}

// GetMustHalt gets the constant fee from the paramSpace
func (k Keeper) GetMustHalt(ctx sdk.Context) (mustHalt bool) {
	k.paramSpace.Get(ctx, types.ParamStoreKeyMustHalt, &mustHalt)

	return
}

// SetMustHalt sets must halt in the paramSpace
func (k Keeper) SetMustHalt(ctx sdk.Context, mustHalt bool) {
	k.paramSpace.Set(ctx, types.ParamStoreKeyMustHalt, mustHalt)
}
