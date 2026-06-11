package main

import "github.com/01-edu/z01"

func IsNegative(chislo int) {
	if chislo < 0 {
		z01.PrintRune('T')
	} else {
		z01.PrintRune('F')
	}
	z01.PrintRune('\n')
}

func main() {
	IsNegative(-1)
	IsNegative(0)
	IsNegative(1)
}
