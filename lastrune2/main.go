package main

import (
	"github.com/01-edu/z01"
)

func LastRune(s string) rune {
	// 1. Превращаем строку в ящик (слайс) с отдельными рунами
	runes := []rune(s)

	// 2. Находим индекс последнего элемента (длина минус 1)
	lastIndex := len(runes) - 1

	// 3. Достаем и возвращаем руну из ящика по этому индексу
	return runes[lastIndex]
}

func main() {
	// Тесты из задания
	z01.PrintRune(LastRune("Hello!"))
	z01.PrintRune(LastRune("Salut!"))
	z01.PrintRune(LastRune("Ola!"))
	z01.PrintRune('\n')
}
