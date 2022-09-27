package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// NewGenesisState creates a new GenesisState object
func NewGenesisState(constantFee sdk.Coin, mustHalt bool) *GenesisState {
	return &GenesisState{
		ConstantFee: constantFee,
		MustHalt:    mustHalt,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		ConstantFee: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)),
		MustHalt:    false,
	}
}

// ValidateGenesis - validate crisis genesis data
func ValidateGenesis(data *GenesisState) error {
	if !data.ConstantFee.IsPositive() {
		return fmt.Errorf("constant fee must be positive: %s", data.ConstantFee)
	}
	return nil
}
