package main

func RectPerimeter(length, width int) int {
	if length < 0 || width < 0 {
		return 0 // Возвращаем 0 для отрицательных значений
	}
	return 2 * (length + width)
}

func main() {
	// Примеры использования функции RectPerimeter
	println(RectPerimeter(5, 10))  // Должно выдать 30
	println(RectPerimeter(0, 10))  // Должно выдать 20
	println(RectPerimeter(-5, 10)) // Должно выдать 0 (отрицательное значение)
}
