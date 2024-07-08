package piscine

func StrLen(s string) uint {
	var i uint
	for range s {
		i++
	}
	return i
}

func ConcatParams(args []string) string {
	var re string
	for _, v := range args {
		re += (v + "\n")
	}
	if StrLen(re) == 0 {
		return re
	}
	return re[:StrLen(re)-1]
}
