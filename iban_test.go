package ibanlib_test

import (
	"log"
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
