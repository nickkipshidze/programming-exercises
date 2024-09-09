package main

import "fmt"

func main() {
	var input string

	fmt.Printf("What is the input string? ")
	fmt.Scan(&input)
	fmt.Printf("%s has %d characters.\n", input, len(input))
}