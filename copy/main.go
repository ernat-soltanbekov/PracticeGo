package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// NbrFunction возвращает: само число, его экспоненту и натуральный логарифм от модуля
func NbrFunction(c int) (int, float64, float64) {
	// Для математики из пакета math нужны дробные числа (float64)
	val := float64(c)

	exponential := math.Exp(val)
	naturalLog := math.Log(math.Abs(val))

	// В Go можно вернуть сразу несколько значений через запятую
	return c, exponential, naturalLog
}

// StrFunction возвращает: исходную строку и строку с экспонентами каждого числа
func StrFunction(a string) (string, string) {
	// Разрезаем строку на отдельные слова по пробелу
	words := strings.Split(a, " ")
	var results []string

	for _, w := range words {
		// Превращаем текст в дробное число
		val, _ := strconv.ParseFloat(w, 64)

		// Вычисляем экспоненту
		expVal := math.Exp(val)

		// Превращаем число обратно в текст (-1 значит "вывести без лишних нулей")
		// и добавляем в новый список
		strExp := strconv.FormatFloat(expVal, 'f', -1, 64)
		results = append(results, strExp)
	}

	// Склеиваем массив обратно в единую строку через пробел
	return a, strings.Join(results, " ")
}

// VecFunction возвращает: исходный массив и массив логарифмов
func VecFunction(b []int) ([]int, []float64) {
	var logs []float64

	for _, val := range b {
		fVal := float64(val)
		// Считаем натуральный логарифм от абсолютного значения (модуля)
		logs = append(logs, math.Log(math.Abs(fVal)))
	}

	return b, logs
}

func main() {
	// Тестовые данные из задания
	c := 0
	b := []int{1, 2, 4, 5}
	a := "1 2 4 5 6"

	// 1. Тест NbrFunction
	valC, expC, logC := NbrFunction(c)
	// В Go логарифм от нуля возвращает отрицательную бесконечность (-Inf)
	fmt.Printf("(%d, %v, %v)\n", valC, expC, logC)

	// 2. Тест VecFunction
	valB, logsB := VecFunction(b)
	fmt.Printf("(%v, %v)\n", valB, logsB)

	// 3. Тест StrFunction
	valA, expA := StrFunction(a)
	fmt.Printf("(\"%s\", \"%s\")\n", valA, expA)
}
