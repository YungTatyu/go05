package piscine

func Strlen(s string) int {
	var i int = 0
	for range s {
		i++
	}
	return i
}

func RuneToDigit(r rune) int {
	if r >= '0' && r <= '9' {
		return int(r - '0')
	} else if r >= 'A' && r <= 'Z' {
		return int(r - 'A' + 10)
	} else if r >= 'a' && r <= 'z' {
		return int(r - 'a' + 10)
	}
	return 0
}

func ToDecimal(nbr string, baseFrom string) int {
	base := Strlen(baseFrom)
	decimal := 0
	for _, v := range nbr {
		digitValue := RuneToDigit(v)
		decimal = decimal*base + digitValue
	}
	return decimal
}

func DecimalToBase(decimal int, base string) string {
	var result string
	baseLen := Strlen(base)
	for decimal > 0 {
		result = string(base[decimal%baseLen]) + result
		decimal /= baseLen
	}
	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return DecimalToBase(ToDecimal(nbr, baseFrom), baseTo)
}
