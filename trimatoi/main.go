package main

import "fmt"

func TrimAtoi(s string) int {
	result := 0         // Наша коробка для сборки итогового числа
	sign := 1           // Знак числа. По умолчанию 1 (положительное)
	foundDigit := false // Флажок: находили ли мы уже хотя бы одну цифру?

	for _, char := range s {
		// Ловим минус. Он работает только если мы еще не находили цифр!
		if char == '-' && !foundDigit {
			sign = -1
		}

		// Если нам попалась цифра
		if char >= '0' && char <= '9' {
			foundDigit = true // Поднимаем флаг: мы начали собирать число

			// Превращаем руну в нормальную цифру и склеиваем с результатом
			digit := int(char - '0')
			result = result*10 + digit
		}
	}

	// Умножаем собранное число на его знак (на 1 или на -1)
	return result * sign
}

func main() {
	// Тесты из твоего README
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}
