package ibanlib

import (
	"fmt"
	"strconv"
	"strings"
)

// Checksum will compute the appropriate checksum value for the current iban value
// and return it.
func (iban *IBAN) Checksum() (int, error) {
	return ibanChecksum(iban.BBAN() + iban.Country + "00")
}

// SetChecksum will update the IBAN's checksum value to ensure it is valid
func (iban *IBAN) SetChecksum() error {
	val, err := iban.Checksum()
	if err != nil {
		return err
	}
	iban.Check = fmt.Sprintf("%02d", val)
	return nil
}

func ibanToNumeric(iban string) (string, error) {
	var numericIBAN strings.Builder

	for _, r := range iban {
		if r >= 'A' && r <= 'Z' {
			numericValue := int(r - 'A' + 10)
			numericIBAN.WriteString(strconv.Itoa(numericValue))
		} else if r >= '0' && r <= '9' {
			numericIBAN.WriteRune(r)
		} else {
			return "", fmt.Errorf("invalid character in IBAN: %c", r)
		}
	}

	return numericIBAN.String(), nil
}

func ibanChecksum(input string) (int, error) {
	numericInput, err := ibanToNumeric(input)
	if err != nil {
		return 0, err
	}

	mod97 := int(0)
	for _, digit := range numericInput {
		num := digit - '0'
		mod97 = (mod97*10 + int(num)) % 97
	}
	return 98 - mod97, nil
}
