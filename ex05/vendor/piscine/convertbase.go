package piscine

import "ft"

func Strlen(s string) uint {
	var i uint = 0
	for range s {
		i++
	}
	return i
}

func Print(s string) {
	for _, v := range s {
		ft.PrintRune(v)
	}
}

func ConvertBase(nbr, baseFrom, baseTo string) string {
	var arr []string

}
