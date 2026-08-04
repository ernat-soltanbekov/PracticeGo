package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("0")
		return
	}

	strVal1 := os.Args[1]
	operator := os.Args[2]
	strVal2 := os.Args[3]

	val1, ok1 := Atoi(strVal1)
	val2, ok2 := Atoi(strVal2)

	if !ok1 || !ok2 {
		return
	}

	switch operator {
	case "+":
		fmt.Println(val1 + val2)
	case "-":
		fmt.Println(val1 - val2)
	case "*":
		fmt.Println(val1 * val2)
	case "/":
		fmt.Println(val1 / val2)
	case "%":
		fmt.Println(val1 % val2)
	default:
		fmt.Println("0")
		return
	}
}

func Atoi(s string) (int, bool) {
	if s == "" {
		return 0, false
	}

	sign := 1
	i := 0
	if s[0] == '-' {
		sign = -1
		i = 1
	} else if s[0] == '+' {
		i = 1
	}

	if i == len(s) {
		return 0, false
	}

	res := 0
	for ; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return 0, false
		}
		digit := int(s[i] - '0')

		if res > 922337203685477580 || (res == 922337203685477580 && digit > 7 && sign == 1) || (res == 922337203685477580 && digit > 8 && sign == -1) {
			return 0, false
		}
		res = res*10 + digit
	}
	return res * sign, true
}
