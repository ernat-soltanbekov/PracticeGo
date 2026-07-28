package main

import "fmt"

func Concat(str1 string, str2 string) string {
	return str1 + str2
}

func main() {
	fmt.Println(Concat("Aku.", " Soku. Zan."))
}
