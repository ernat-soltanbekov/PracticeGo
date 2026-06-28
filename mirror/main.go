package main

import "fmt"

func main() {
	// Исходная переменная, как указано в задании
	word := "desserts" //

	// 1. Превращаем строку в массив (срез) рун
	runes := []rune(word)

	// 2. Алгоритм переворота (Reverse)
	// i идет от 0 (начала), а j идет от конца массива.
	// Цикл работает, пока i меньше j (пока они не встретятся в центре).
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		// Фишка языка Go: меняем две переменные местами без создания третьей!
		runes[i], runes[j] = runes[j], runes[i]
	}

	// 3. Собираем массив обратно в строку и объявляем переменную mirror
	mirror := string(runes) //

	// Выводим результат в терминал Ubuntu для проверки
	fmt.Println("Original word:", word)
	fmt.Println("Mirrored word:", mirror)
}
