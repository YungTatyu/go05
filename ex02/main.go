package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.ConcatParams([]string{"Hello", "how", "are", "you?"}))
	fmt.Println(piscine.ConcatParams([]string{"test"}))
	fmt.Println(piscine.ConcatParams([]string{""}))
	fmt.Println(piscine.ConcatParams([]string{"", ""}))
}
