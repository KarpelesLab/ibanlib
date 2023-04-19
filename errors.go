package ibanlib

import "errors"

var (
	ErrTooShort             = errors.New("IBAN too short")
	ErrTooLong              = errors.New("IBAN too long")
	ErrInvalidCountry       = errors.New("invalid or unknown IBAN country")
	ErrInvalidChecksum      = errors.New("invalid IBAN checksum")
	ErrInvalidBankLength    = errors.New("invalid IBAN bank id length")
	ErrInvalidBranchLength  = errors.New("invalid IBAN branch id length")
	ErrInvalidAccountLength = errors.New("invalid IBAN account id length")
)
