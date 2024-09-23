package main

import "fmt"

func main() {
	var amount float32
	var state string

	fmt.Printf("What is the order amount? ")
	fmt.Scan(&amount)

	fmt.Printf("What is the state? ")
	fmt.Scan(&state)

	if state == "WI" {
		tax := amount * 0.055
		fmt.Printf("The subtotal is $%.2f.\n", amount)
		fmt.Printf("The tax is $%.2f.\n", tax)
		amount += tax
	}

	fmt.Printf("The total is $%.2f.\n", amount)
}