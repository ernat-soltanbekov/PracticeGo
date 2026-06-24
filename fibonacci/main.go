package main

import "fmt"

func Fibonacci(index int) int {
	// 1. Базовый случай: защита от отрицательных чисел
	if index < 0 {
		return -1
	}

	// 2. Базовый случай: начало последовательности
	if index == 0 {
		return 0
	}

	// 3. Базовый случай: шаг номер 1
	if index == 1 {
		return 1
	}

	// 4. Рекурсивный шаг (циклов for нет!)
	// Функция вызывает саму себя ДВАЖДЫ, чтобы узнать два предыдущих числа,
	// складывает их и возвращает результат наверх.
	return Fibonacci(index-1) + Fibonacci(index-2)
}

func main() {
	// Тест из твоего README
	arg1 := 4
	fmt.Println(Fibonacci(arg1)) // Ожидаем 3, так как последовательность: 0, 1, 1, 2, [3]

	// Дополнительные тесты для уверенности
	fmt.Println(Fibonacci(0))  // Ожидаем 0
	fmt.Println(Fibonacci(-5)) // Ожидаем -1
	fmt.Println(Fibonacci(7))  // Ожидаем 13
}
