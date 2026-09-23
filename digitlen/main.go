package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.DigitLen(100, 10))  // Должно выдать 3
	fmt.Println(piscine.DigitLen(100, 2))   // Должно выдать 7
	fmt.Println(piscine.DigitLen(-100, 16)) // Должно выдать 2
	fmt.Println(piscine.DigitLen(100, -1))  // Должно выдать -1
}
