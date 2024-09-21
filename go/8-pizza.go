package main

import "fmt"

func main() {
	var people, pizzas, slices int

	fmt.Printf("How many people? ")
	fmt.Scan(&people)

	fmt.Printf("How many pizzas do you have? ")
	fmt.Scan(&pizzas)

	fmt.Printf("How may slices per pizza? ")
	fmt.Scan(&slices)

	fmt.Printf("%d people with %d pizzas\n", people, pizzas)
    fmt.Printf("Each person gets %d pieces of pizza.\n", (pizzas * slices) / people)
    fmt.Printf("There are %d leftover pieces.\n", (pizzas * slices) % people)
}