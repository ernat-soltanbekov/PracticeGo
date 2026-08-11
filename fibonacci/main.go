package main

import "fmt"

func Fibonacci(index int) int {
	if index == 0 {
		return 0
	}
	if index == 1 {
		return 1
	}
	result := Fibonacci(index-2) + Fibonacci(index-1)
	return result
}

func main() {
	fmt.Println(Fibonacci(7))
}
