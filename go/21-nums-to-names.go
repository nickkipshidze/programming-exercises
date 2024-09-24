package main

import "fmt"

func main() {
	var number int
	var month string

	fmt.Printf("Please enter the number of the month: ")
	fmt.Scan(&number)

	switch number {
		case 1: month = "January"
		case 2: month = "February"
		case 3: month = "March"
		case 4: month = "April"
		case 5: month = "May"
		case 6: month = "June"
		case 7: month = "July"
		case 8: month = "August"
		case 9: month = "September"
		case 10: month = "October"
		case 11: month = "November"
		case 12: month = "December"
	}

	fmt.Printf("The name of the month is %s.\n", month)
}