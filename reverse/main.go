package main

import "fmt"

// Структура узла
type NodeAddL struct {
	Next *NodeAddL
	Num  int
}

// ТВОЯ ФУНКЦИЯ
func Reverse(node *NodeAddL) *NodeAddL {
	// Пиши логику здесь
	return nil
}

// Вспомогательная функция, чтобы main мог собрать список
func pushBack(n *NodeAddL, num int) *NodeAddL {
	newNode := &NodeAddL{Num: num}
	if n == nil {
		return newNode
	}
	current := n
	for current.Next != nil {
		current = current.Next
	}
	current.Next = newNode
	return n
}

func main() {
	num1 := &NodeAddL{Num: 1}
	num1 = pushBack(num1, 3)
	num1 = pushBack(num1, 2)
	num1 = pushBack(num1, 4)
	num1 = pushBack(num1, 5)

	result := Reverse(num1)

	// Вывод результата
	for tmp := result; tmp != nil; tmp = tmp.Next {
		fmt.Print(tmp.Num)
		if tmp.Next != nil {
			fmt.Print(" -> ")
		}
	}
	fmt.Println()
}
