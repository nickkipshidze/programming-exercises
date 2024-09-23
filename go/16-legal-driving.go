package main

import "fmt"

func main() {
	var age int

	fmt.Printf("What is your age? ")
	fmt.Scan(&age)

	if age < 16 {
		fmt.Printf("You are not old enough to legally drive.\n")
	} else {
		fmt.Printf("You are old enough to legally drive.\n")
	}
}