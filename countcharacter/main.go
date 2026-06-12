package main

import (
	"fmt"
)

func CountChar(str string, c rune) int {
	count := 0
	for _, character := range str {
		if character == c {
			count++
		}
	}
	return count
}

func main() {
	fmt.Println(CountChar("Hello World", 'o'))
	fmt.Println(CountChar("5  balloons", 98))
	fmt.Println(CountChar("nursultan", 'u'))
	fmt.Println(CountChar("The 7 deadly sins", ' '))
}
