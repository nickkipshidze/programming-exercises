package main

import "fmt"

func main() {
	var first, second, third, largest int

	fmt.Printf("Enter the first number: ")
	fmt.Scan(&first)
	fmt.Printf("Enter the second number: ")
	fmt.Scan(&second)
	fmt.Printf("Enter the third number: ")
	fmt.Scan(&third)

	if first > second && first > third {
		largest = first
	}
	if second > first && second > third {
		largest = second
	}
	if third > second && third > first {
		largest = third
	}

	fmt.Printf("The largest number is %d.\n", largest)
}