package main //

import "fmt" //

// NodeAddL — структура нашего узла, согласно инструкции
type NodeAddL struct { //[cite: 5]
	Next *NodeAddL //[cite: 5]
	Num  int       //[cite: 5]
} //[cite: 5]

// Reverse разворачивает список и возвращает новый корень[cite: 5]
func Reverse(node *NodeAddL) *NodeAddL { //[cite: 5]
	var prev *NodeAddL = nil // Предыдущий узел (изначально пусто)
	current := node          // Текущий узел (начинаем с переданной головы)

	// Пока мы не дойдем до конца поезда
	for current != nil {
		// 1. Страховка: запоминаем остаток поезда
		next := current.Next

		// 2. Перепайка: разворачиваем сцепку текущего вагона назад
		current.Next = prev

		// 3. Шаг вперед: сдвигаем "предыдущий" на место текущего
		prev = current

		// 4. Шаг вперед: сдвигаем "текущий" на сохраненный остаток поезда
		current = next
	}

	// Когда цикл закончится, current укажет на пустоту (nil),
	// а prev будет указывать на бывший последний вагон. Он и есть наша новая голова!
	return prev
}

// pushBack добавляет новый узел в самый конец списка (нужно для тестов)[cite: 5]
func pushBack(n *NodeAddL, num int) *NodeAddL { //[cite: 5]
	// Если список вообще пустой, создаем первый вагон
	if n == nil {
		return &NodeAddL{Num: num}
	}

	// Идем по списку до самого конца
	current := n
	for current.Next != nil {
		current = current.Next
	}

	// Прицепляем новый вагон в хвост
	current.Next = &NodeAddL{Num: num}
	return n
}

func main() { //[cite: 5]
	// Сборка тестового списка из задания: 1 -> 3 -> 2 -> 4 -> 5[cite: 5]
	num1 := &NodeAddL{Num: 1} //[cite: 5]
	num1 = pushBack(num1, 3)  //[cite: 5]
	num1 = pushBack(num1, 2)  //[cite: 5]
	num1 = pushBack(num1, 4)  //[cite: 5]
	num1 = pushBack(num1, 5)  //[cite: 5]

	// Вызываем нашу функцию переворота[cite: 5]
	result := Reverse(num1) //[cite: 5]

	// Вывод результата в консоль[cite: 5]
	for tmp := result; tmp != nil; tmp = tmp.Next { //[cite: 5]
		fmt.Print(tmp.Num)   //[cite: 5]
		if tmp.Next != nil { //[cite: 5]
			fmt.Print(" -> ") //[cite: 5]
		} //[cite: 5]
	} //[cite: 5]
	fmt.Println() //[cite: 5]
} //[cite: 5]
