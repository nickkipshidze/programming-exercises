package main

import "fmt"

func main() {
	const taxrate float32 = 0.055
	var price, quantity [3]int
	var subtotal int = 0
	
	for i := 0; i < 3; i++ {
		fmt.Printf("Enter the price of an item %d: ", i+1)
		fmt.Scan(&price[i])
		fmt.Printf("Enter the quantity of an item %d: ", i+1)
		fmt.Scan(&quantity[i])

		subtotal += price[i] * quantity[i]
	}

	tax := float32(subtotal) * taxrate
	total := float32(subtotal) + tax

	fmt.Printf("Subtotal: $%.2f\nTax: $%.2f\nTotal: $%.2f\n", float32(subtotal), tax, total)
}