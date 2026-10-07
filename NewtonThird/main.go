package main

import "fmt"

func NewtonThird(force float64) float64 {
	return -force
}

func main() {
	action := 3359.0
	reaction := NewtonThird(action)

	fmt.Println("Action:", action, "N")
	fmt.Println("Reaction:", reaction, "N")
}
