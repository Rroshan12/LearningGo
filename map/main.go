package main

import "fmt"

func main() {
	fmt.Println("Learning map")
	studentGrades := make(map[string]int)
	studentGrades["Roshan"] = 98
	studentGrades["Hari"] = 100
	studentGrades["Ram"] = 100
	fmt.Println(studentGrades)
	fmt.Println("Roshan marks", studentGrades["Roshan"])
	grades, exists := studentGrades["David"]
	fmt.Println(grades, exists)
	delete(studentGrades, "Hari")
	fmt.Println(studentGrades)

	for index, value := range studentGrades {
		fmt.Println(index, ":", value)
	}

	numbername := map[string]int{
		"Ramesh": 1,
		"Rohit":  2,
	}

	for index, value := range numbername {
		fmt.Println(index, ":", value)
	}
}
