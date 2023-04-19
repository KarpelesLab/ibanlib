package ibanlib

import "math/rand"

const (
	AlphaRange        = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	NumericRange      = "0123456789"
	AlphanumericRange = AlphaRange + NumericRange
)

// RandomValue return a low quality random value suitable for account numbers etc
func RandomValue(acceptableRange string, length int) string {
	// generate a random string of the required length
	b := make([]rune, length)
	// prepare input
	input := []rune(acceptableRange)
	inputLn := len(input)
	// build
	for i := range b {
		b[i] = input[rand.Intn(inputLn)]
	}
	return string(b)
}
