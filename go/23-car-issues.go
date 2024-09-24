package main

import "fmt"

func prompt(message string) (bool) {
	var input string

	fmt.Printf("%s", message)
	fmt.Scan(&input)

	if input == "y" {
		return true
	} else {
		return false
	}
}

func main() {
	if prompt("Is the car silent when you turn the key? ") {
		if prompt("Are the battery terminals corroded? ") {
			fmt.Printf("Clean the terminals and try starting again.\n")
		} else {
			fmt.Printf("Replace cables and try again.\n")
		}
	} else {
		if prompt("Does the car make a clicking noise? ") {
			fmt.Printf("Replace the battery.\n")
		} else {
			if prompt("Does the car crank up but fail to start? ") {
				fmt.Printf("Check the spark plug connections.\n")
			} else {
				if prompt("Does the engine start and then die? ") {
					if prompt("Does your car have fuel injection? ") {
						fmt.Printf("Get it in for service.\n")
					} else {
						fmt.Printf("Check to ensure the choke is opening and closing.\n")
					}
				} else {
					fmt.Print("Get outta here then!\n")
				}
			}
		}
	}
}