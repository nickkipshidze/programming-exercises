package main

import "fmt"

func msquared(fsquared int) float32 {
	return float32(fsquared) * 0.09290304
}

func main() {
	var lenft, widthft int

	fmt.Printf("What is the length of the room in feet? ")
	fmt.Scan(&lenft)

	fmt.Printf("What is the width of the room in feet? ")
	fmt.Scan(&widthft)

	fmt.Printf("You entered dimensions of %d feet by %d feet.\n", lenft, widthft)

	var squarefeet int = lenft * widthft

	squaremeters := msquared(squarefeet)

	fmt.Printf("The area is\n%d square feet\n%f square meters\n", squarefeet, squaremeters)
}