package main

import (
	"fmt"
)

func main() {

	day := 3

	month := "March"

	switch day {
	case 1:
		fmt.Println("One")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("one")
	}

	switch month {
	case "January", "March":
		fmt.Println("Nice")
	default:
		fmt.Println("Good")
	}

	switch {
	case day > 1:
		fmt.Println("Greater One")
	case day >2:
		fmt.Println("Gretaer Three")
	default:
		fmt.Println("one")
	}
}
