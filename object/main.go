package main

import "fmt"

// 1. Создаем "чертеж" для нашего объекта
// В Go принято называть поля структуры с большой буквы
type Person struct {
	Name          string
	Age           int
	SecureLuggage bool
}

func main() {
	// 2. Объявляем переменную human, как того требует инструкция
	human := Person{
		Name:          "GhostOfAstana", // Текстовое значение в кавычках
		Age:           32,              // Обычное целое число
		SecureLuggage: true,            // true (безопасно) или false (опасно)
	}

	// 3. Выводим результат в терминал
	fmt.Println(human)
}
