package main

import "fmt"

func main() {

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	count := 0
	for {

		if count > 3 {
			break
		}
		fmt.Println("Infinite Loop")
		count++
	}

	numbers := [] int{1,2,3,4}

	for index, value := range numbers{
		fmt.Println(index, value)
	}

}
