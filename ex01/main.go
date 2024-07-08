package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.MakeRange(0, 0))
	fmt.Println(piscine.MakeRange(0, -1))
	fmt.Println(piscine.MakeRange(5, 10))
	fmt.Println(piscine.MakeRange(1, 100))
}
