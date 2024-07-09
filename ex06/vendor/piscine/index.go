package piscine

func Strlen(s string) uint {
	var i uint = 0
	for range s {
		i++
	}
	return i
}

func StrnCmp(s1 string, s2 string, len uint) int {
	var i uint = 0
	s1Len := Strlen(s1)
	s2Len := Strlen(s2)
	for i < len {
		if i >= s1Len || i >= s2Len {
			return -1
		}
		var diff int = int(s1[i] - s2[i])
		if diff != 0 {
			return diff
		}
		i++
	}
	return 0
}

func Index(s string, toFind string) int {
	if toFind == "" {
		return -1
	}
	var i int = 0
	for range s {
		if StrnCmp(s[i:], toFind, Strlen(toFind)) == 0 {
			return i
		}
		i++
	}
	return -1
}
