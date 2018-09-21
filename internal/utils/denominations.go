package clefutils

import (
	"math/big"
)

const (
	Wei   = 1
	GWei  = 1e9
	Ether = 1e18
)

func ConvertUnit(value *big.Float, fromUnit float64, toUnit float64) *big.Float {
	if fromUnit < toUnit {
		ratio := big.NewFloat(0).Quo(big.NewFloat(toUnit), big.NewFloat(fromUnit))
		return big.NewFloat(0).Quo(value, ratio)
	}

	if fromUnit > toUnit {
		ratio := big.NewFloat(0).Quo(big.NewFloat(fromUnit), big.NewFloat(toUnit))
		return big.NewFloat(0).Mul(value, ratio)
	}

	return value
}

func ConvertUnitAndGetString(value *big.Float, fromUnit float64, toUnit float64) string {
	floatValue := ConvertUnit(value, fromUnit, toUnit)
	return floatValue.Text(byte('f'), -1)
}