package clefutils

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	Wei   = uint64(1)
	GWei  = uint64(1e9)
	Ether = uint64(1e18)
)

// TextToWei converts a string representation of a certain unit into a big.Int (in wei)
func TextToWei(text string, fromUnit uint64) (*big.Int, error) {
	valueFloat, _, err := big.ParseFloat(text, 10, 0, big.ToNearestEven);
	if err != nil {
		return nil, fmt.Errorf("failed to parse float value %s", text)
	}
	// Convert unit into wei
	valueFloat = valueFloat.Mul(valueFloat, big.NewFloat(0).SetUint64(fromUnit))
	bigV,_ := valueFloat.Int(new(big.Int))
	return bigV, nil
}

// WeiToString converts bigint wei into a string-representation in the given unit
func WeiToString(weiVal *big.Int, toUnit uint64) string {
	floatV := new(big.Float).SetInt(weiVal)
	floatV.Quo(floatV, new(big.Float).SetUint64(toUnit))
	return fmt.Sprintf("%g", floatV)
}

// DefaultFormat converts bigint wei into a non-exact but user-friendly form
func DefaultFormat(weiVal *big.Int) string {
	var unit uint64
	var unitStr string
	// Show in Gwei if above 0.001 Gwei
	if weiVal.Cmp(big.NewInt(1e6)) < 0{
		unit = Wei
		unitStr = "wei"
	}else if weiVal.Cmp(big.NewInt(1e15)) < 0{
		unit = GWei
		unitStr = "Gwei"
	}else{
		unit = Ether
		unitStr = "eth"
	}
	floatV := new(big.Float).SetInt(weiVal)
	floatV.Quo(floatV, new(big.Float).SetUint64(unit))
	return fmt.Sprintf("%.3g %s", floatV, unitStr)
}

func ToChecksumAddress(rawAddress string) (returnAddress string, err error) {
	a, err := common.NewMixedcaseAddressFromString(rawAddress)
	if err != nil{
		return "", err
	}
	return a.Address().String(), nil
}

func IsValidAddress(rawAddress string) bool {
	a, err := common.NewMixedcaseAddressFromString(rawAddress)
	if err != nil{
		return false
	}
	return a.ValidChecksum()
}
