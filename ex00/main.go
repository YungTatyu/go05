package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.AppendRange(0, 0))
	fmt.Println(piscine.AppendRange(0, -1))
	fmt.Println(piscine.AppendRange(5, 10))
	fmt.Println(piscine.AppendRange(1, 100))
}
