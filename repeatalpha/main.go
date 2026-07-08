package main

import (
	"fmt"
)

// RepeatAlpha принимает строку и повторяет символы согласно их номеру в алфавите
func RepeatAlpha(s string) string { //
	// Создаем пустой массив рун для сборки нового слова
	var result []rune

	// Перебираем каждый символ в исходной строке
	for _, char := range s {
		repeatCount := 1 // По умолчанию печатаем символ 1 раз (для пробелов, цифр и т.д.)

		// Если это строчная буква
		if char >= 'a' && char <= 'z' {
			// Вычисляем её порядковый номер в алфавите (от 1 до 26)
			repeatCount = int(char-'a') + 1
		} else if char >= 'A' && char <= 'Z' { // Если это заглавная буква
			// Делаем то же самое, но вычитаем заглавную 'A'
			repeatCount = int(char-'A') + 1
		}

		// Добавляем текущий символ в наш массив repeatCount раз
		for i := 0; i < repeatCount; i++ {
			result = append(result, char)
		}
	}

	// Превращаем собранный массив рун обратно в единую строку и возвращаем
	return string(result)
}

func main() { //
	// Тестовые данные из инструкции
	fmt.Println(RepeatAlpha("abc"))           //
	fmt.Println(RepeatAlpha("Choumi."))       //
	fmt.Println(RepeatAlpha(""))              //
	fmt.Println(RepeatAlpha("abacadaba 01!")) //
}
