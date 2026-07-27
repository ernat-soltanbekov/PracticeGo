package main

import "fmt"

func Concat(str1 string, str2 string, str3 string) string {
	return str1 + str2 + str3
}

func main() {
	fmt.Println(Concat("Aku.", " Soku.", " Zan."))
}
