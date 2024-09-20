package main

import "fmt"

const YEAR int = 2015

func main() {
	var age, retire int

	fmt.Printf("What is your current age? ")
	fmt.Scan(&age)

	fmt.Printf("At what age would you like to retire? ")
	fmt.Scan(&retire)

	remaining := retire - age
	
	fmt.Printf("You have %d years left until you can retire.\n", remaining)
	fmt.Printf("It's %d, so you can retire in %d.\n", YEAR, YEAR + remaining)
}