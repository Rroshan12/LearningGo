package main

import "fmt"

func main() {
	fmt.Println("Starting dynamic array slice")
	num := []int{1, 2, 3, 4}
	fmt.Printf("%T \n", num)
	fmt.Println(num)
	num = append(num, 4, 5, 6)
	fmt.Println(num, len(num), cap(num))

	// make function empty slice
	numbers := make([]int, 3, 5)
	emptyslice := make([]int, 0)
	fmt.Println(numbers, len(numbers), cap(numbers), emptyslice)
	numbers = append(numbers, 4, 5, 6)
	fmt.Println(numbers, len(numbers), cap(numbers))

}
