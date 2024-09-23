package main

import "fmt"

func main() {
	const password string = "abc$123"
	var buffer string

	fmt.Printf("What is the password? ")
	fmt.Scan(&buffer)

	if buffer == password {
		fmt.Printf("Welcome!\n")
	} else {
		fmt.Printf("I don't know you.\n")
	}
}