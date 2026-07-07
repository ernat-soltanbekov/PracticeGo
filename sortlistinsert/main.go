package main

import "fmt"

// NodeI — структура нашего узла
type NodeI struct {
	Data int
	Next *NodeI
}

// SortListInsert вставляет новый элемент, сохраняя порядок сортировки
func SortListInsert(l *NodeI, data_ref int) *NodeI {
	// 1. Создаем новый узел в памяти с переданным значением
	newNode := &NodeI{Data: data_ref}

	// Сценарий 1: Если изначальный список абсолютно пуст
	if l == nil {
		return newNode
	}

	// Сценарий 2: Новое значение меньше или равно значению ПЕРВОГО элемента (головы).
	// Значит, наш новый узел должен стать новой головой списка.
	if data_ref <= l.Data {
		newNode.Next = l // Привязываем старую цепочку позади нового узла
		return newNode   // Возвращаем новый узел как новое начало
	}

	// Сценарий 3: Ищем правильное место в середине или конце списка.
	current := l

	// Двигаемся вперед по списку.
	// Условие работы цикла: следующий вагон существует И значение в следующем вагоне меньше нашего.
	// Как только следующий вагон окажется больше нашего (или вагоны закончатся), цикл остановится.
	for current.Next != nil && current.Next.Data < data_ref {
		current = current.Next
	}

	// Мы нашли нужное место! current указывает на вагон, ПОСЛЕ которого нужно вставить новый.
	// Сначала прицепляем к новому вагону ту часть состава, которая идет после current
	newNode.Next = current.Next
	// А затем говорим current, что теперь за ним идет наш новый вагон
	current.Next = newNode

	// Возвращаем оригинальную голову списка (так как она не изменилась)
	return l
}

// --- Вспомогательные функции для проверки (go run .) ---

func PrintList(l *NodeI) { //
	it := l
	for it != nil {
		fmt.Print(it.Data, " -> ")
		it = it.Next
	}
	fmt.Print(nil, "\n")
}

func listPushBack(l *NodeI, data int) *NodeI { //
	n := &NodeI{Data: data}

	if l == nil {
		return n
	}
	iterator := l
	for iterator.Next != nil {
		iterator = iterator.Next
	}
	iterator.Next = n
	return l
}

func main() { //
	var link *NodeI

	// Формируем начальный отсортированный список: 1 -> 4 -> 9
	link = listPushBack(link, 1) //
	link = listPushBack(link, 4) //
	link = listPushBack(link, 9) //

	PrintList(link) //

	// Вставляем элементы и проверяем результат
	link = SortListInsert(link, -2) //
	link = SortListInsert(link, 2)  //
	PrintList(link)                 //
}
