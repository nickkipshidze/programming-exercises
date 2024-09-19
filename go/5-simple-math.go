package main

import "fmt"

func main() {
	var first, second int
	
	fmt.Printf("What is the first number? ")
	fmt.Scan(&first)
	fmt.Printf("What is the second number? ")
	fmt.Scan(&second)

	fmt.Printf("%d + %d = %d\n", first, second, first + second)
	fmt.Printf("%d - %d = %d\n", first, second, first - second)
	fmt.Printf("%d * %d = %d\n", first, second, first * second)
	fmt.Printf("%d / %d = %d\n", first, second, first / second)
}