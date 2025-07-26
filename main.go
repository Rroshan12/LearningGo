package main

import (
	"fmt"
	"mylearning/myutil"
)

func main() {
	fmt.Println("hjj")
	myutil.PrintMessage("Hello World")
	var name string = "Roshan"
	var num int = 10
	var price float64 = 86.3
	check()
	fmt.Println(name, num, price)
}

func check() {
	fmt.Print("Apple check")
}
