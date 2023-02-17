package types

import (
	// "math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	ethermint "github.com/evmos/ethermint/types"
)

var (
	// total supply of bond coin
	// g_totalSupply = sdk.NewInt(int64(700000000000000))
	g_totalSupply = sdk.TokensFromConsensusPower(int64(700000000), ethermint.PowerReduction)

	// inflation of a new block
	// 3 * (60 * 60 * 8760 / 6) = sdk.NewDec(15768000000000)
	// g_inflationPerBlock = sdk.NewInt(1200000)
	g_inflationPerBlock = sdk.TokensFromConsensusPower(12, ethermint.PowerReduction.QuoRaw(10))

	// interval times
	g_everyYears = int64(2)
)

type MintHook struct {
}

func NewMintHook() *MintHook {
	return &MintHook{}
}

func (h MintHook) BeforeNextAnnualProvisions(blockHeight int64, blocksPerYear uint64, totalSupply sdk.Int, customProvision bool) sdk.Dec {
	rc := sdk.NewDec(-1)

	if(totalSupply.GT(g_totalSupply)){
		rc = sdk.ZeroDec()
	}else if(customProvision){ 
		// g_inflationPerBlock / (2 ^ (currentBlockHeight / (years * numberOfPerYear)))
		// exponent := new(big.Int).Quo(big.NewInt(blockHeight), big.NewInt(int64(blocksPerYear) * g_everyYears))
		// fractor := sdk.NewIntFromBigInt(new(big.Int).Exp(big.NewInt(2), exponent, nil))
		// inflationPerBlock := g_inflationPerBlock.Quo(fractor)
		// rc = sdk.NewDecFromInt(inflationPerBlock).MulInt(sdk.NewInt(int64(blocksPerYear)))
		rc = sdk.NewDecFromInt(g_inflationPerBlock).MulInt(sdk.NewInt(int64(blocksPerYear)))
	}

	return rc
}
