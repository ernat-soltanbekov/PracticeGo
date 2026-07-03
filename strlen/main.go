package main

import "fmt"

// StrLen принимает строку и возвращает количество рун (символов) в ней
func StrLen(s string) int { //
	// Конвертируем сырую строку (байты) в массив полноценных символов (рун)
	runes := []rune(s)

	// Возвращаем количество элементов в получившемся массиве
	return len(runes)
}

func main() {
	// Тестовый пример из инструкции
	l := StrLen("Hello World!") //
	fmt.Println(l)              //
}
