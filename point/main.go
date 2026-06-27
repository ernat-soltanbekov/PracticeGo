package main

import "fmt"

// Создаем наш собственный тип данных 'point'
type point struct {
	x int
	y int
}

// Функция принимает УКАЗАТЕЛЬ на структуру point
func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	// Создаем пустую структуру point и сразу берем ее адрес в памяти (&)
	points := &point{}

	// Передаем адрес структуры в функцию
	setPoint(points)

	// Распечатываем результат с помощью форматирования
	fmt.Printf("x = %d, y = %d\n", points.x, points.y)
}
