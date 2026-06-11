package main

import "fmt"

func ToUpper(basarip string) string {
	basaripter := []rune(basarip)
	for ULKEN := 0; ULKEN < len(basaripter); ULKEN++ {
		if basaripter[ULKEN] >= 'a' && basaripter[ULKEN] <= 'z' {
			basaripter[ULKEN] -= 32
		}
	}
	return string(basaripter)
}

func main() {
	fmt.Println(ToUpper("hello"))
}
