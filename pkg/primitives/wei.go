package primitives

import "math/big"

func WeiToETH(wei *big.Int) float64 {
	decimals := 1e18
	weiFloat := new(big.Float).SetInt(wei)
	decimalsBigFloat := big.NewFloat(decimals)
	ethValue := new(big.Float).Quo(weiFloat, decimalsBigFloat)
	eth, _ := ethValue.Float64()

	return eth
}
