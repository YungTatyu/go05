package piscine

func ToDecimal(nbr string, baseFrom string) int {
	base := int(Strlen(baseFrom))
	decimal := 0
	for _, v := range nbr {
		digitValue := Index(baseFrom, string(v))
		decimal = decimal*base + digitValue
	}
	return decimal
}

func DecimalToBase(decimal int, base string) string {
	var result string
	baseLen := int(Strlen(base))
	for decimal > 0 {
		result = string(base[decimal%baseLen]) + result
		decimal /= baseLen
	}
	return result
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	return DecimalToBase(ToDecimal(nbr, baseFrom), baseTo)
}
