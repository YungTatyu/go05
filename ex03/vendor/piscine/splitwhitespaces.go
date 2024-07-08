package piscine

const (
	StateStart = iota
	StateSpace
	StateWord
)

func StrLen(s string) uint {
	var i uint
	for range s {
		i++
	}
	return i
}

func IsSpace(r rune) bool {
	return r == ' ' || r == '\n' || r == '\t'
}

func SplitWhiteSpaces(s string) []string {
	var i uint
	len := StrLen(s)
	var arr []string
	state := StateStart
	var curValue string
	for {
		if i >= len {
			break
		}
		r := rune(s[i])
		switch state {
		case StateStart:
			state = StateSpace
		case StateSpace:
			if !IsSpace(r) {
				state = StateWord
				break
			}
			i++
		case StateWord:
			if IsSpace(r) {
				arr = append(arr, curValue)
				state = StateSpace
				curValue = ""
				i++
				break
			}
			curValue += string(r)
			i++
		}
	}
	if curValue != "" {
		arr = append(arr, curValue)
	}
	return arr
}
