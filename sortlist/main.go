package main

import (
	"fmt"
)

// NodeList — структура нашего узла, как требует инструкция
type NodeList struct { //
	Data int       //
	Next *NodeList //
}

// SortList сортирует связный список
func SortList(l *NodeList, cmp func(a, b int) bool) *NodeList { //[cite: 10]
	// Защита: если список пустой или в нем только один элемент, он уже отсортирован
	if l == nil || l.Next == nil {
		return l
	}

	swapped := true
	// Крутим цикл до тех пор, пока происходят обмены
	for swapped {
		swapped = false
		current := l

		// Проходим по цепочке до предпоследнего элемента
		for current.Next != nil {
			// Передаем значения в функцию сравнения
			if !cmp(current.Data, current.Next.Data) {
				// Если порядок неверный, меняем Data местами
				current.Data, current.Next.Data = current.Next.Data, current.Data
				// Фиксируем, что замена произошла, значит нужен еще один проход
				swapped = true
			}
			// Шагаем дальше по цепочке
			current = current.Next
		}
	}

	// Возвращаем указатель на первый элемент отсортированного списка
	return l //[cite: 10]
}

// ascending - тестовая функция сравнения из инструкции
func ascending(a, b int) bool { //[cite: 10]
	if a <= b { //[cite: 10]
		return true //[cite: 10]
	} else { //[cite: 10]
		return false //[cite: 10]
	}
}

func main() {
	// Создаем тестовый список с конца в начало: 5 -> 2 -> 4 -> 1 -> 3
	node5 := &NodeList{Data: 3, Next: nil}
	node4 := &NodeList{Data: 1, Next: node5}
	node3 := &NodeList{Data: 4, Next: node4}
	node2 := &NodeList{Data: 2, Next: node3}
	head := &NodeList{Data: 5, Next: node2}

	fmt.Print("Оригинальный список: ")
	for temp := head; temp != nil; temp = temp.Next {
		fmt.Print(temp.Data, " ")
	}
	fmt.Println()

	// Запускаем сортировку, передавая саму функцию ascending как аргумент
	head = SortList(head, ascending)

	fmt.Print("Отсортированный список: ")
	for temp := head; temp != nil; temp = temp.Next {
		fmt.Print(temp.Data, " ")
	}
	fmt.Println()
}
