package main

import (
	"fmt"
)

func oFunction() {
	fmt.Println("oooooooo")
}

func add(a, b int) int {
	return a + b
}

func sub(a, b int) (result int) {
	result = a - b
	return result
}

func main() {
	fmt.Println("we are learning function")
	oFunction()
	sub := sub(3, 1)
	fmt.Println(add(1, 2), sub)
}
