package main

import "fmt"

func main() {
	num := 1
	ptr := &num
	fmt.Println(ptr, *ptr)
	value := 3
	modifyValueByRef(&value)
	fmt.Println(value)

}

func modifyValueByRef(num *int) {
	*num = *num * 2
}
