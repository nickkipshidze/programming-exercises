package main

import "fmt"

func exchange(amountFrom, rateFrom, rateTo float32) (float32) {
	return (amountFrom * rateFrom) / rateTo
}

func main() {
	var euros, exchangeRate float32

	fmt.Printf("How many euros are you exchanging? ")
	fmt.Scan(&euros)
	fmt.Printf("What is the exchange rate? ")
	fmt.Scan(&exchangeRate)

	dollars := exchange(euros, exchangeRate, 100.36)

	fmt.Printf("%d euros at an exchange rate of %.2f is %.2f U.S. dollars.\n", int(euros), exchangeRate, dollars)
}