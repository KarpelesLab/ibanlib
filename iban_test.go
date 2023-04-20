package ibanlib_test

import (
	"log"
	"math/rand"
	"testing"

	"github.com/KarpelesLab/ibanlib"
)

func TestIbanParsing(t *testing.T) {
	// IBAN belonging to the Red Cross: https://www.icrc.org/en/faq/donate-by-bank-transfer-or-cheque
	ib, err := ibanlib.Parse("CH97 0024 0240 FP10 0883 2")
	log.Printf("ib = %+v", *ib)
	if err != nil {
		t.Errorf("error while parsing valid iban: %s", err)
	}
}

func TestIbanRandom(t *testing.T) {
	// generate random account
	ib := &ibanlib.IBAN{Country: "FR", Bank: "00001", Branch: "00001"}
	rand.Seed(42)
	ib.SetRandomAccount()
	if ib.String() != "FR56 0000 1000 0157 8035 7683 975" {
		t.Errorf("error while generating random iban, got %s", ib)
	}
}
