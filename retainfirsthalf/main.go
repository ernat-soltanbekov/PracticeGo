package main

import "fmt"

func RetainFirstHalf(str string) string {
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	half := len(str) / 2
	return str[:half]
}

func main() {
	fmt.Println(RetainFirstHalf("Last Checkpoint"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
