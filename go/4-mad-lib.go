package main

import "fmt"

func main() {
	var noun, verb, adjective, adverb string

	fmt.Printf("Enter a noun: ")
	fmt.Scan(&noun)
	fmt.Printf("Enter a verb: ")
	fmt.Scan(&verb)
	fmt.Printf("Enter an adjective: ")
	fmt.Scan(&adjective)
	fmt.Printf("Enter an adverb: ")
	fmt.Scan(&adverb)

	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adjective, noun, adverb)
}