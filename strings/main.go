package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "apple,orange,banana"
	parts := strings.Split(data, ",")
	fmt.Println(parts)
	str := "one two three two one five six seven four five six four"
	count := strings.Count(str, "one")
	fmt.Println(count)
}
