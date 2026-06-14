package main

import "fmt"

func LastWord(s string) string {
	lastWord := ""    // Сейф для последнего сохраненного слова
	currentWord := "" // Черновик для сборки текущего слова по буквам

	for _, char := range s {
		if char == ' ' {
			// Мы встретили пробел.
			// Если в черновике что-то есть, значит слово только что закончилось.
			if currentWord != "" {
				lastWord = currentWord // Сохраняем собранное слово в сейф
				currentWord = ""       // Очищаем черновик для следующего слова
			}
		} else {
			// Если это не пробел (буква/знак), приклеиваем в черновик
			currentWord += string(char)
		}
	}

	// ЛОВУШКА: Строка могла закончиться буквой, а не пробелом.
	// Если мы дошли до конца строки, а в черновике осталось недосохраненное слово,
	// мы обязаны переложить его в сейф, иначе оно потеряется.
	if currentWord != "" {
		lastWord = currentWord
	}

	// Возвращаем результат + перенос строки (если слов не было, lastWord пустой)
	return lastWord + "\n"
}

func main() {
	// Тесты из твоего README для проверки
	fmt.Print(LastWord("this        ...       is sparta, then again, maybe    not"))
	fmt.Print(LastWord(" lorem,ipsum "))
	fmt.Print(LastWord(" "))
}
