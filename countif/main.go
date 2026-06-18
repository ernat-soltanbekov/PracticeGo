package main

import "fmt"

func IsNumeric(s string) bool {
	for _, char := range s {
		if char < '0' || char > '9' {
			return false
		}
	}
	return true
}

func CountIf(f func(string) bool, tab []string) int {
	count := 0
	for _, str := range tab {
		if f(str) {
			count++
		}
	}
	return count
}

func main() {
	tab1 := []string{"123", "abc", "456", "def", "789"}
	tab2 := []string{"hello", "world", "go", "programming"}

	fmt.Println(CountIf(IsNumeric, tab1))
	fmt.Println(CountIf(IsNumeric, tab2))
}
