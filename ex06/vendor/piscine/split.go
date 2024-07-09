package piscine

func Split(s, sep string) []string {
	if s == "" {
		return nil
	}
	var arr []string
	sepLen := Strlen(sep)
	for {
		re := Index(s, sep)
		if re == -1 {
			break
		}
		if re != 0 {
			arr = append(arr, s[:re])
		}
		s = s[uint(re)+sepLen:]
	}
	if s != "" {
		arr = append(arr, s)
	}
	return arr
}
