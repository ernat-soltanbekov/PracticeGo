package main

import "fmt"

func IsPrime(nb int) bool {
	if nb <= 1 {
		return false
	}
	for i := 2; i*i <= nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("5:", IsPrime(5))     // true
	fmt.Println("4:", IsPrime(4))     // false
	fmt.Println("1:", IsPrime(1))     // false
	fmt.Println("97:", IsPrime(97))   // true (большое простое)
	fmt.Println("100:", IsPrime(100)) // false
}
