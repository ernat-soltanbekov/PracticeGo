package main

import "fmt"

func FirstWord(s string) string {
	result := "" // Наш аккумулятор для первого слова

	for _, char := range s {
		if char == ' ' {
			// Если мы встретили пробел, проверяем коробку.
			// Если в result уже есть буквы, значит слово собрано. Тормозим цикл!
			if result != "" {
				break
			}
			// Если result пустой, значит это пробел в самом начале строки.
			// Мы просто ничего не делаем, цикл идет дальше к следующему символу.
		} else {
			// Если это не пробел, превращаем руну в строку и приклеиваем
			result += string(char)
		}
	}

	// Когда цикл закончился (сам или через break), приклеиваем перенос строки
	return result + "\n"
}

func main() {
	// Тесты из задания
	fmt.Print(FirstWord("hello there"))
	fmt.Print(FirstWord(""))
	fmt.Print(FirstWord("hello   .........  bye"))

	// Дополнительный тест: что если пробелы стоят в самом начале?
	fmt.Print(FirstWord("   testword  "))
}
