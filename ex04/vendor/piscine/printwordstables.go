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

func PrintWordsTables(a []string) {
	var s string
	for _, v := range a {
		s += (v + "\n")
	}
	if Strlen(s) == 0 {
		return
	}
	Print(s[:Strlen(s)-1])
}
