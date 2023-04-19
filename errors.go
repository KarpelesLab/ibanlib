package ibanlib

import "errors"

var (
	ErrTooShort       = errors.New("IBAN too short")
	ErrTooLong        = errors.New("IBAN too long")
	ErrInvalidCountry = errors.New("invalid or unknown IBAN country")
)
