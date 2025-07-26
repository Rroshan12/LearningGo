package main

import "fmt"

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("Error")
	}
	return a / b, nil
}

func main() {
	fmt.Println("Error handling")
	ans, err := divide(10, 0)
	fmt.Println("Ans: ", ans, err)
}
