package main

import (
	"fmt"
)

func main() {
	fmt.Println("hi, this is aico")
	var name string
	fmt.Println("who are you?")
	fmt.Scanln(&name)
	fmt.Println("nice to meet you,", name)
}
