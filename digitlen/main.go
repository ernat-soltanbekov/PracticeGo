package main

import (
	"fmt"
)

func DigitLen(n int, base int) int {
	if base < 2 || base > 36 {
		return -1
	}
	if n == 0 {
		return 1
	}
	count := 0
	if n < 0 {
		n = -n
	}
	for n > 0 {
		n /= base
		count++
	}
	return count
}

func main() {
	fmt.Println(DigitLen(100, 10))  // Должно выдать 3
	fmt.Println(DigitLen(100, 2))   // Должно выдать 7
	fmt.Println(DigitLen(-100, 16)) // Должно выдать 2
	fmt.Println(DigitLen(100, -1))  // Должно выдать -1
}
