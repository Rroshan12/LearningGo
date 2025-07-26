package main

import (
	"fmt"
	"strconv"
)

func main() {

	var num int = 42

	var data float64 = float64(num)
	fmt.Printf("Type od data %T\n", data)
	str := strconv.Itoa(num)
	fmt.Printf("%q\n",str)
	fmt.Printf("%T\n", str)

}
