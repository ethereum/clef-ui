package clefutils

import (
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
	"regexp"
	"strings"
)

// Errors
var (
	ErrOddLength      = &decError{"hex string of odd length"}
	ErrInvalidAddress = &decError{"address is invalid"}
)

var AddressPattern = regexp.MustCompile(`(?i)(^0x)?([0-9]|[A-Z]){40}`)

type decError struct{ msg string }

func (err decError) Error() string { return err.msg }

func ToChecksumAddress(rawAddress string) (returnAddress string, err error) {
	matched := AddressPattern.MatchString(rawAddress)

	if len(rawAddress) != 42 {
		return "", ErrOddLength
	}

	if !matched {
		fmt.Println("Doesn't match pattern")
		return "", ErrInvalidAddress
	}

	address := []byte(strings.ToLower(rawAddress[2:]))

	h := sha3.NewLegacyKeccak256()
	h.Write(address)
	hash := h.Sum(nil)
	hashStrings := hex.EncodeToString(hash[:])

	ret := "0x"
	for i, charByte := range address {
		hashInt := decodeNibble(hashStrings[i])

		if hashInt > 8 {
			ret += strings.ToUpper(string(charByte))
		} else {
			ret += string(charByte)
		}
	}

	return ret, nil
}

func IsValidAddress(rawAddress string) bool {
	ret, err := ToChecksumAddress(rawAddress)

	if err != nil {
		return false
	}

	return ret == rawAddress
}

func decodeNibble(in byte) uint64 {
	if in <= 55 {
		return 0
	} else {
		return 9
	}
}
