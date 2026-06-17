package main

import "fmt"

func Concat(str1 string, str2 string) string {
	return str1 + str2
}

func main() {
	fmt.Println(Concat("Hello, ", "world!"))
	fmt.Println(Concat("Go ", "programming"))
	fmt.Println(Concat("Concatenation ", "is fun!"))
}
