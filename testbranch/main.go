package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Intn(100) // Generate a random number between 0 and 99
	fmt.Println("Random Number:", randomNumber)
}
