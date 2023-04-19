package ibanlib

import (
	"bytes"
	"fmt"
	"strings"
)

type IBAN struct {
	Country string // 2 letters
	Check   string // 2 digits
	Bank    string // bank code
	Branch  string // branch or sort code
	Account string // account number
}

// Parse will parse a given IBAN and always return a value, even if an error happens
func Parse(iban string) (*IBAN, error) {
	// remove all non alphanumeric chars
	iban = cleanIban(iban)

	if len(iban) < 4 {
		return &IBAN{Account: iban}, ErrTooShort
	}
	res := &IBAN{
		Country: strings.ToUpper(iban[0:2]),
		Check:   iban[2:4],
	}

	rule, ok := rules[res.Country]
	if !ok {
		res.Account = iban
		return res, ErrInvalidCountry
	}
	// check length
	if len(iban) < rule.total {
		res.Account = iban
		return res, ErrTooShort
	}
	if len(iban) > rule.total {
		res.Account = iban
		return res, ErrTooLong
	}
	iban = iban[4:]

	// length is confirmed, set fields
	res.Bank = iban[:rule.bank]
	iban = iban[rule.bank:]
	if rule.branch > 0 {
		res.Branch = iban[:rule.branch]
		iban = iban[rule.branch:]
	}
	res.Account = iban

	return res, res.IsValid()
}

// IsValid will return nil if the iban is valid
func (iban *IBAN) IsValid() error {
	ckSum, err := iban.Checksum()
	if err != nil {
		return err
	}
	if iban.Check != fmt.Sprintf("%02d", ckSum) {
		return ErrInvalidChecksum
	}

	rule, ok := rules[iban.Country]
	if !ok {
		return ErrInvalidCountry
	}

	if len(iban.Bank) != rule.bank {
		return ErrInvalidBankLength
	}
	if len(iban.Branch) != rule.branch {
		return ErrInvalidBankLength
	}
	if len(iban.Account) != rule.account {
		return ErrInvalidBankLength
	}
	return nil
}

// CompactString returns a IBAN without any spaces/etc
func (iban *IBAN) CompactString() string {
	str := iban.Country + iban.Check + iban.Bank + iban.Branch + iban.Account
	return str
}

// BBAN returns the Basic Bank Account Number part of the IBAN
func (iban *IBAN) BBAN() string {
	return iban.Bank + iban.Branch + iban.Account
}

// String returns a formatted IBAN
func (iban *IBAN) String() string {
	str := iban.CompactString()
	res := &bytes.Buffer{}
	first := true

	for {
		if len(str) > 4 {
			if !first {
				res.WriteByte(' ')
			}
			res.WriteString(str[:4])
			str = str[4:]
			first = false
			continue
		}
		if len(str) > 0 {
			if !first {
				res.WriteByte(' ')
			}
			res.WriteString(str)
		}
		break
	}

	return res.String()
}
