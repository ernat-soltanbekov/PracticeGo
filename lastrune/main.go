package main

func LastRune(s string) rune {
	var last rune
	for _, char := range s {
		last = char
	}
	return last
}

func main() {
	println(LastRune("Aka Soku Zan"))
	println(LastRune("Ola!"))
	println(LastRune(""))
}
