package keeper

import (
	paramtypes "github.com/elys-network/elys/x/parameter/types"
)

func (k Keeper) GetNoneNativeAsset(collateralAsset string, borrowAsset string) string {
	if collateralAsset == paramtypes.USDC {
		return borrowAsset
	}

	return collateralAsset
}
