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

	fmt.Println("\naico can do a lot of things")

	fmt.Print("\n")

	fmt.Println("what do want aico to do?")
	fmt.Println("1.  list all the files and folders in the current directory")
	fmt.Println("2.  search for a file")
	fmt.Println("3.  create a file")
	fmt.Println("4.  delete a file")
	fmt.Println("5.  rename a file")
	fmt.Println("6.  duplicate a file")
	fmt.Println("7.  move a file")
	fmt.Println("8.  list all the files and folders [ including hidden ones ] in the current directory")
	fmt.Println("9.  search for a directory")
	fmt.Println("10. create a directory")
	fmt.Println("11. delete a directory")
	fmt.Println("12. rename a directory")
	fmt.Println("13. duplicate a directory")
	fmt.Println("14. move a directory")

	fmt.Print("\n")
}
