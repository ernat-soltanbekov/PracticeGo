package main

import "fmt"

func RecursiveFactorial(nb int) int {
	if nb < 0 || nb > 20 { // Проверка на отрицательные числа и переполнение
		return 0
	} else if nb == 0 {
		return 1
	} else {
		return nb * RecursiveFactorial(nb-1)
	}
}

func main() {
	fmt.Println(RecursiveFactorial(4))  // Должно выдать 24
	fmt.Println(RecursiveFactorial(0))  // Должно выдать 1
	fmt.Println(RecursiveFactorial(-5)) // Должно выдать 0
	fmt.Println(RecursiveFactorial(25)) // Должно выдать 0 (overflow)
}
