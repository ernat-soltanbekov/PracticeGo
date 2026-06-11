package main

import "fmt"

func PrintVowels(s string) {
	for _, c := range s {
		if c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U' {
			fmt.Printf("%c!", c)
		}
	}
	fmt.Println()
}

func main() {
	PrintVowels("Aka. Soku. Zan.")
}
