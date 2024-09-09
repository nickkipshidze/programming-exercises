package main

import "fmt"

func main() {
	var quote, author string

	fmt.Printf("What is the quote? ")
	fmt.Scanln(&quote)
	fmt.Printf("Who said it? ")
	fmt.Scanln(&author)

	fmt.Printf("%s says, \"%s\"\n", author, quote)
}