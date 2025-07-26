package main

import "fmt"

func main() {

	fmt.Println("Starting array")

	var name [5]string

	var numbers = [8]int{1, 2, 3, 4, 5}

	name[0] = "roshan"
	name[1] = "saleena"

	fmt.Printf("%q", numbers)

	fmt.Println(name, numbers, len(numbers))

}
