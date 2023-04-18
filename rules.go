package ibanlib

type ibanRule struct {
	// rules for a specific country
	bank    int  // bank id length
	branch  int  // branch id length (0 if none)
	natid   int  // national id length, typically bank+branch
	account int  // account number length
	total   int  // total length, should be == bank+branch+account+4
	sepa    bool // is this a SEPA IBAN?
}

// TODO check & update rules
var rules = map[string]*ibanRule{
	"AD": &ibanRule{bank: 4, branch: 4, natid: 8, account: 12, total: 24},
	"AE": &ibanRule{bank: 3, branch: 0, natid: 3, account: 16, total: 23},
	"AL": &ibanRule{bank: 3, branch: 5, natid: 8, account: 16, total: 28},
	"AT": &ibanRule{bank: 5, branch: 0, natid: 5, account: 11, total: 20, sepa: true},
	"AZ": &ibanRule{bank: 4, branch: 0, natid: 4, account: 20, total: 28},
	"BA": &ibanRule{bank: 3, branch: 3, natid: 3, account: 10, total: 20},
	"BE": &ibanRule{bank: 3, branch: 0, natid: 3, account: 9, total: 16, sepa: true},
	"BG": &ibanRule{bank: 4, branch: 4, natid: 8, account: 10, total: 22, sepa: true},
	"BH": &ibanRule{bank: 4, branch: 0, natid: 4, account: 14, total: 22},
	"BR": &ibanRule{bank: 8, branch: 5, natid: 8, account: 12, total: 29},
	"CH": &ibanRule{bank: 5, branch: 0, natid: 5, account: 12, total: 21, sepa: true},
	"CR": &ibanRule{bank: 3, branch: 0, natid: 3, account: 14, total: 21},
	"CY": &ibanRule{bank: 3, branch: 5, natid: 8, account: 16, total: 28, sepa: true},
	"CZ": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24, sepa: true},
	"DE": &ibanRule{bank: 8, branch: 0, natid: 8, account: 10, total: 22, sepa: true},
	"DK": &ibanRule{bank: 4, branch: 0, natid: 4, account: 10, total: 18, sepa: true},
	"DO": &ibanRule{bank: 4, branch: 0, natid: 4, account: 20, total: 28},
	"EE": &ibanRule{bank: 2, branch: 0, natid: 2, account: 14, total: 20, sepa: true},
	"ES": &ibanRule{bank: 4, branch: 4, natid: 8, account: 12, total: 24, sepa: true},
	"FI": &ibanRule{bank: 6, branch: 0, natid: 3, account: 8, total: 18, sepa: true},
	"FO": &ibanRule{bank: 4, branch: 0, natid: 4, account: 10, total: 18},
	"FR": &ibanRule{bank: 5, branch: 5, natid: 10, account: 13, total: 27, sepa: true},
	"GB": &ibanRule{bank: 4, branch: 6, natid: 10, account: 8, total: 22, sepa: true},
	"GE": &ibanRule{bank: 2, branch: 0, natid: 2, account: 16, total: 22},
	"GI": &ibanRule{bank: 4, branch: 0, natid: 4, account: 15, total: 23, sepa: true},
	"GL": &ibanRule{bank: 4, branch: 0, natid: 4, account: 10, total: 18},
	"GR": &ibanRule{bank: 3, branch: 4, natid: 7, account: 16, total: 27, sepa: true},
	"GT": &ibanRule{bank: 4, branch: 0, natid: 4, account: 20, total: 28},
	"HR": &ibanRule{bank: 7, branch: 0, natid: 7, account: 10, total: 21, sepa: true},
	"HU": &ibanRule{bank: 3, branch: 4, natid: 7, account: 17, total: 28, sepa: true},
	"IE": &ibanRule{bank: 4, branch: 6, natid: 10, account: 8, total: 22, sepa: true},
	"IL": &ibanRule{bank: 3, branch: 3, natid: 6, account: 13, total: 23},
	"IS": &ibanRule{bank: 4, branch: 0, natid: 4, account: 18, total: 26, sepa: true},
	"IT": &ibanRule{bank: 5, branch: 5, natid: 10, account: 12, total: 27, sepa: true},
	"KW": &ibanRule{bank: 4, branch: 0, natid: 4, account: 22, total: 30},
	"KZ": &ibanRule{bank: 3, branch: 0, natid: 3, account: 13, total: 20},
	"LB": &ibanRule{bank: 4, branch: 0, natid: 4, account: 14, total: 28},
	"LI": &ibanRule{bank: 5, branch: 0, natid: 5, account: 12, total: 21, sepa: true},
	"LT": &ibanRule{bank: 5, branch: 0, natid: 5, account: 11, total: 20, sepa: true},
	"LU": &ibanRule{bank: 3, branch: 0, natid: 3, account: 13, total: 20, sepa: true},
	"LV": &ibanRule{bank: 4, branch: 0, natid: 4, account: 13, total: 21, sepa: true},
	"MC": &ibanRule{bank: 5, branch: 5, natid: 10, account: 13, total: 27, sepa: true},
	"MD": &ibanRule{bank: 2, branch: 0, natid: 2, account: 18, total: 24},
	"ME": &ibanRule{bank: 3, branch: 0, natid: 3, account: 15, total: 22},
	"MK": &ibanRule{bank: 3, branch: 0, natid: 3, account: 12, total: 19},
	"MR": &ibanRule{bank: 5, branch: 5, natid: 10, account: 13, total: 27},
	"MT": &ibanRule{bank: 4, branch: 5, natid: 9, account: 18, total: 31, sepa: true},
	"MU": &ibanRule{bank: 6, branch: 2, natid: 8, account: 18, total: 30},
	"NL": &ibanRule{bank: 4, branch: 0, natid: 4, account: 10, total: 18, sepa: true},
	"NO": &ibanRule{bank: 4, branch: 0, natid: 4, account: 7, total: 15, sepa: true},
	"PK": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24},
	"PL": &ibanRule{bank: 8, branch: 0, natid: 8, account: 16, total: 28, sepa: true},
	"PS": &ibanRule{bank: 4, branch: 0, natid: 4, account: 21, total: 29},
	"PT": &ibanRule{bank: 4, branch: 4, natid: 8, account: 13, total: 25, sepa: true},
	"RO": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24, sepa: true},
	"RS": &ibanRule{bank: 3, branch: 0, natid: 3, account: 15, total: 22},
	"SA": &ibanRule{bank: 2, branch: 0, natid: 2, account: 18, total: 24},
	"SE": &ibanRule{bank: 3, branch: 0, natid: 3, account: 17, total: 24, sepa: true},
	"SI": &ibanRule{bank: 5, branch: 0, natid: 2, account: 10, total: 19, sepa: true},
	"SK": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24, sepa: true},
	"SM": &ibanRule{bank: 5, branch: 5, natid: 10, account: 12, total: 27},
	"TN": &ibanRule{bank: 2, branch: 3, natid: 2, account: 15, total: 24},
	"TR": &ibanRule{bank: 5, branch: 0, natid: 5, account: 17, total: 26},
	"VG": &ibanRule{bank: 4, branch: 0, natid: 4, account: 16, total: 24},
}
