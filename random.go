package ibanlib

import (
	"math/rand"
	"sync"
	"time"
)

const (
	AlphaRange        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumericRange      = "0123456789"
	AlphanumericRange = AlphaRange + NumericRange
)

var (
	rnd     *rand.Rand
	rndOnce sync.Once
)

// SetRandomSeed must not need to be used in normal operations, and is only useful
// to have tests that can be reproduced. By default we will use UnixNano() to initialize
// our random seed, so you really don't need to call this.
func SetRandomSeed(n int64) {
	// spend the once
	rndOnce.Do(func() {})
	rnd = rand.New(rand.NewSource(n))
}

// RandomValue return a low quality random value suitable for account numbers etc
func RandomValue(acceptableRange string, length int) string {
	rndOnce.Do(func() {
		rnd = rand.New(rand.NewSource(time.Now().UnixNano()))
	})
	// generate a random string of the required length
	b := make([]rune, length)
	// prepare input
	input := []rune(acceptableRange)
	inputLn := len(input)
	// build
	for i := range b {
		b[i] = input[rnd.Intn(inputLn)]
	}
	return string(b)
}

func (iban *IBAN) SetRandomAccount() error {
	// generate random account
	rule, ok := rules[iban.Country]
	if !ok {
		return ErrInvalidCountry
	}

	// TODO apply other rules, etc
	iban.Account = RandomValue(NumericRange, rule.account)
	return iban.SetChecksum()
}
