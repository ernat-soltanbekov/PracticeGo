package main

import "fmt"

func IsNegative(a int) bool {
	if a < 0 {
		return true
	}
	return false
}

func main() {
	number := -5

	result := IsNegative(number)

	fmt.Println(result)
}
