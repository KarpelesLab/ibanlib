package ibanlib

type IBAN struct {
	Country string // 2 letters
	Bank    string // bank code
	Branch  string // branch or sort code
	Account string // account number
}
