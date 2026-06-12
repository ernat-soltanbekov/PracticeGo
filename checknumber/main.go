package main

import "fmt"

func CheckNumber(arg string) bool {
//	for i := 0; i < len(arg); i++ {
//		if arg[i] >= '0' && arg[i] <= '9' {
//			return true
//		}
//	}
//	return false
//}

    for _, char := range arg {
        if char >= '0' && char <= '9' {
            return true
        }
    }
    return false
}

func main() {
    fmt.Println(CheckNumber("Hello"))
    fmt.Println(CheckNumber("Hello1"))
}
