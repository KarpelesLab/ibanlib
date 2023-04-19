package ibanlib

type ibanRule struct {
	// rules for a specific country
	bank    int    // bank id length
	branch  int    // branch id length (0 if none)
	natid   int    // national id length, typically bank+branch
	account int    // account number length
	total   int    // total length, should be == bank+branch+account+4
	sepa    bool   // is this a SEPA IBAN?
	format  string // expected format
}

// TODO check & update rules
var rules = map[string]*ibanRule{
	"AD": &ibanRule{bank: 4, branch: 4, natid: 8, account: 12, total: 24, format: "8n,12c"},
	"AE": &ibanRule{bank: 3, branch: 0, natid: 3, account: 16, total: 23, format: "3n,16n"},
	"AL": &ibanRule{bank: 3, branch: 5, natid: 8, account: 16, total: 28, format: "8n,16c"},
	"AT": &ibanRule{bank: 5, branch: 0, natid: 5, account: 11, total: 20, sepa: true, format: "16n"},
	"AZ": &ibanRule{bank: 4, branch: 0, natid: 4, account: 20, total: 28},
	"BA": &ibanRule{bank: 3, branch: 3, natid: 3, account: 10, total: 20, format: "16n"},
	"BE": &ibanRule{bank: 3, branch: 0, natid: 3, account: 9, total: 16, sepa: true, format: "12n"},
	"BG": &ibanRule{bank: 4, branch: 4, natid: 8, account: 10, total: 22, sepa: true, format: "4a,6n,8c"},
	"BH": &ibanRule{bank: 4, branch: 0, natid: 4, account: 14, total: 22},
	"BR": &ibanRule{bank: 8, branch: 5, natid: 8, account: 12, total: 29},
	"CH": &ibanRule{bank: 5, branch: 0, natid: 5, account: 12, total: 21, sepa: true, format: "5n,12c"},
	"CR": &ibanRule{bank: 3, branch: 0, natid: 3, account: 14, total: 21},
	"CY": &ibanRule{bank: 3, branch: 5, natid: 8, account: 16, total: 28, sepa: true, format: "8n,16c"},
	"CZ": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24, sepa: true, format: "20n"},
	"DE": &ibanRule{bank: 8, branch: 0, natid: 8, account: 10, total: 22, sepa: true, format: "18n"},
	"DK": &ibanRule{bank: 4, branch: 0, natid: 4, account: 10, total: 18, sepa: true, format: "14n"},
	"DO": &ibanRule{bank: 4, branch: 0, natid: 4, account: 20, total: 28, format: "4a,20n"},
	"EE": &ibanRule{bank: 2, branch: 0, natid: 2, account: 14, total: 20, sepa: true, format: "16n"},
	"ES": &ibanRule{bank: 4, branch: 4, natid: 8, account: 12, total: 24, sepa: true, format: "20n"},
	"FI": &ibanRule{bank: 6, branch: 0, natid: 3, account: 8, total: 18, sepa: true, format: "14n"},
	"FO": &ibanRule{bank: 4, branch: 0, natid: 4, account: 10, total: 18, format: "14n"},
	"FR": &ibanRule{bank: 5, branch: 5, natid: 10, account: 13, total: 27, sepa: true, format: "10n,11c,2n"},
	"GB": &ibanRule{bank: 4, branch: 6, natid: 10, account: 8, total: 22, sepa: true, format: "4a,14n"},
	"GE": &ibanRule{bank: 2, branch: 0, natid: 2, account: 16, total: 22, format: "2c,16n"},
	"GI": &ibanRule{bank: 4, branch: 0, natid: 4, account: 15, total: 23, sepa: true, format: "4a,15c"},
	"GL": &ibanRule{bank: 4, branch: 0, natid: 4, account: 10, total: 18, format: "14n"},
	"GR": &ibanRule{bank: 3, branch: 4, natid: 7, account: 16, total: 27, sepa: true, format: "7n,16c"},
	"GT": &ibanRule{bank: 4, branch: 0, natid: 4, account: 20, total: 28},
	"HR": &ibanRule{bank: 7, branch: 0, natid: 7, account: 10, total: 21, sepa: true, format: "17n"},
	"HU": &ibanRule{bank: 3, branch: 4, natid: 7, account: 17, total: 28, sepa: true, format: "24n"},
	"IE": &ibanRule{bank: 4, branch: 6, natid: 10, account: 8, total: 22, sepa: true, format: "4c,14n"},
	"IL": &ibanRule{bank: 3, branch: 3, natid: 6, account: 13, total: 23, format: "19n"},
	"IS": &ibanRule{bank: 4, branch: 0, natid: 4, account: 18, total: 26, sepa: true, format: "22n"},
	"IT": &ibanRule{bank: 5, branch: 5, natid: 10, account: 12, total: 27, sepa: true, format: "1a,10n,12c"},
	"KW": &ibanRule{bank: 4, branch: 0, natid: 4, account: 22, total: 30, format: "4a,22n"},
	"KZ": &ibanRule{bank: 3, branch: 0, natid: 3, account: 13, total: 20, format: "3n,3c,10c"},
	"LB": &ibanRule{bank: 4, branch: 0, natid: 4, account: 14, total: 28, format: "4n,20c"},
	"LI": &ibanRule{bank: 5, branch: 0, natid: 5, account: 12, total: 21, sepa: true, format: "5n,12c"},
	"LT": &ibanRule{bank: 5, branch: 0, natid: 5, account: 11, total: 20, sepa: true, format: "16n"},
	"LU": &ibanRule{bank: 3, branch: 0, natid: 3, account: 13, total: 20, sepa: true, format: "3n,13c"},
	"LV": &ibanRule{bank: 4, branch: 0, natid: 4, account: 13, total: 21, sepa: true, format: "4a,13c"},
	"MC": &ibanRule{bank: 5, branch: 5, natid: 10, account: 13, total: 27, sepa: true, format: "10n,11c,2n"},
	"MD": &ibanRule{bank: 2, branch: 0, natid: 2, account: 18, total: 24},
	"ME": &ibanRule{bank: 3, branch: 0, natid: 3, account: 15, total: 22, format: "18n"},
	"MK": &ibanRule{bank: 3, branch: 0, natid: 3, account: 12, total: 19, format: "3n,10c,2n"},
	"MR": &ibanRule{bank: 5, branch: 5, natid: 10, account: 13, total: 27, format: "23n"},
	"MT": &ibanRule{bank: 4, branch: 5, natid: 9, account: 18, total: 31, sepa: true, format: "4a,5n,18c"},
	"MU": &ibanRule{bank: 6, branch: 2, natid: 8, account: 18, total: 30, format: "4a,19n,3a"},
	"NL": &ibanRule{bank: 4, branch: 0, natid: 4, account: 10, total: 18, sepa: true, format: "4a,10n"},
	"NO": &ibanRule{bank: 4, branch: 0, natid: 4, account: 7, total: 15, sepa: true, format: "11n"},
	"PK": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24},
	"PL": &ibanRule{bank: 8, branch: 0, natid: 8, account: 16, total: 28, sepa: true, format: "24n"},
	"PS": &ibanRule{bank: 4, branch: 0, natid: 4, account: 21, total: 29},
	"PT": &ibanRule{bank: 4, branch: 4, natid: 8, account: 13, total: 25, sepa: true, format: "21n"},
	"RO": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24, sepa: true, format: "4a,16c"},
	"RS": &ibanRule{bank: 3, branch: 0, natid: 3, account: 15, total: 22, format: "18n"},
	"SA": &ibanRule{bank: 2, branch: 0, natid: 2, account: 18, total: 24, format: "2n,18c"},
	"SE": &ibanRule{bank: 3, branch: 0, natid: 3, account: 17, total: 24, sepa: true, format: "20n"},
	"SI": &ibanRule{bank: 5, branch: 0, natid: 2, account: 10, total: 19, sepa: true, format: "15n"},
	"SK": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24, sepa: true, format: "20n"},
	"SM": &ibanRule{bank: 5, branch: 5, natid: 10, account: 12, total: 27, format: "1a,10n,12c"},
	"TN": &ibanRule{bank: 2, branch: 3, natid: 2, account: 15, total: 24, format: "20n"},
	"TR": &ibanRule{bank: 5, branch: 0, natid: 5, account: 17, total: 26, format: "5n,17c"},
	"VG": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24},
}

// Accounts in those countries will start with the provided code for IBAN
var countryAlias = map[string]string{
	"PF": "FR", // Polynesie Francaise
	"TF": "FR",
	"YT": "FR",
	"NC": "FR",
	"PM": "FR",
	"WF": "FR",
}
