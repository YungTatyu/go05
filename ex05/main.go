package main

import (
	"piscine"
)

func main() {
	piscine.PrintWordsTables([]string{"hello", "how", "are", "you"})
	piscine.PrintWordsTables([]string{"test"})
	piscine.PrintWordsTables([]string{""})
}
