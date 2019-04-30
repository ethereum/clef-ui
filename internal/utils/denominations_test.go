package clefutils

import (
	"math/big"
	"strings"
	"testing"
)

func TestTextToBig(t *testing.T) {

	for i, tt := range []struct {
		input string
		unit  uint64
		exp   *big.Int
	}{
		{"1", Wei, big.NewInt(1)},
		{"1", GWei, big.NewInt(1000000000)},
		{"1", Ether, new(big.Int).Mul(big.NewInt(1000000000), big.NewInt(1000000000))},
		{"1.5", GWei, big.NewInt(1500000000)},
		{"0.1", GWei, big.NewInt(100000000)},
		{"0.0", Wei, big.NewInt(0)},
		{"1.1", Wei, big.NewInt(1)},
		{"2.91e-07", GWei, big.NewInt(291)},
		// Expected errors
		{"johnny", Wei, nil},
		{"1 wei", Wei, nil},
	} {
		got, err := TextToWei(tt.input, tt.unit)

		if tt.exp == nil && err != nil {
			continue
		}

		if tt.exp == nil && err == nil {
			t.Errorf("test %d: expected err, got none (%s -> %v)", i, tt.input, got)
			continue
		} else if tt.exp != nil && err != nil {
			t.Errorf("test %d: expected no err, got %v (%s)", i, err, tt.input)
			continue
		}

		if tt.exp.Cmp(got) != 0 {
			t.Errorf("test %d: got %v expected %v", i, got, tt.exp)
		}
	}
}

func TestBigToText(t *testing.T) {

	for i, tt := range []struct {
		weival *big.Int
		unit   uint64
		exp    string
	}{
		{big.NewInt(1), Wei, "1"},
		{big.NewInt(1000000000), GWei, "1"},
		{big.NewInt(1987654321), GWei, "1.987654321"},
		{new(big.Int).Mul(big.NewInt(1000000000), big.NewInt(1000000000)), Ether, "1"},
		{big.NewInt(1500000000), GWei, "1.5"},
		{big.NewInt(100000000), GWei, "0.1"},
		{big.NewInt(0), Wei, "0"},
		{big.NewInt(1), Wei, "1"},
		{big.NewInt(291), GWei, "2.91e-07"},
	} {
		got := WeiToString(tt.weival, tt.unit)

		if strings.Compare(got, tt.exp) != 0 {
			t.Errorf("test %d: got %v expected %v", i, got, tt.exp)
		}

	}
}

func TestWeiToDefault(t *testing.T) {

	for i, tt := range []struct {
		weival *big.Int
		exp    string
	}{
		{big.NewInt(1), "1 wei"},
		{big.NewInt(1000000000), "1 Gwei"},
		{big.NewInt(1987654321), "1.99 Gwei"},
		{new(big.Int).Mul(big.NewInt(1000000000), big.NewInt(1000000000)), "1 eth"},
		{big.NewInt(1500000000), "1.5 Gwei"},
		{big.NewInt(100000000), "0.1 Gwei"},
		{big.NewInt(0), "0 wei"},
		{big.NewInt(1), "1 wei"},
		{big.NewInt(291), "291 wei"},
		{big.NewInt(2910), "2.91e+03 wei"},
		{big.NewInt(29100), "2.91e+04 wei"},
		{big.NewInt(291000), "2.91e+05 wei"},
		{big.NewInt(2910000), "0.00291 Gwei"},
		{big.NewInt(29100000), "0.0291 Gwei"},
		{big.NewInt(291000000), "0.291 Gwei"},
		{big.NewInt(2910000000), "2.91 Gwei"},
		{big.NewInt(29100000000), "29.1 Gwei"},
	} {
		got := DefaultFormat(tt.weival)

		if strings.Compare(got, tt.exp) != 0 {
			t.Errorf("test %d: got %v expected %v", i, got, tt.exp)
		}

	}
}
