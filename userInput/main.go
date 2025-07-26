package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hey what is your name")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	fmt.Println("Hello,", name)
}
