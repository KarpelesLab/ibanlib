package ibanlib

func cleanIban(iban string) string {
	// remove any non-alphanumeric character
	s := make([]byte, 0, len(iban))

	for _, b := range iban {
		switch {
		case 'a' <= b && b <= 'z':
			s = append(s, byte(b)-32) // make it uppercase
		case ('A' <= b && b <= 'Z') || ('0' <= b && b <= '9'):
			s = append(s, byte(b))
		}
	}
	return string(s)
}
