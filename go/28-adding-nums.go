package main

import "fmt"

func main() {
	var num, sum int = 0, 0

	for i := 0; i < 5; i++ {
		fmt.Printf("Enter a number: ")
		fmt.Scan(&num)
		sum += num
	}

	fmt.Printf("The total is %d.\n", sum)
}