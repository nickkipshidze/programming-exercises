package main

import (
	"fmt"
	"unicode"
)

func btoi(b bool) (int) {
	if b {
		return 1
	} else {
		return 0
	}
}

func validateFirstName(firstName string) (bool) {
	if firstName == "" {
		fmt.Printf("The first name must be filled in.\n")
		return false
	}
	if len(firstName) <= 3 {
		fmt.Printf("\"%s\" is not a valid first name. It is too short.\n", firstName)
		return false
	}
	return true
}

func validateLastName(lastName string) (bool) {
	if lastName == "" {
		fmt.Printf("The last name must be filled in.\n")
		return false
	}
	if len(lastName) <= 3 {
		fmt.Printf("\"%s\" is not a valid last name. It is too short.\n", lastName)
		return false
	}
	return true
}

func validateZipCode(zipCode string) (bool) {
	for _, char := range zipCode {
		if !unicode.IsDigit(char) {
			fmt.Printf("The ZIP code must be numeric.\n")
			return false
		}
	}
	return true
}

func validateEmployeeId(eid string) (bool) {
	// Whatever man. If it works - it works!
	if len(eid) == 7 && unicode.IsLetter(rune(eid[0])) && unicode.IsLetter(rune(eid[1])) && string(eid[2]) == "-" && unicode.IsDigit(rune(eid[3])) && unicode.IsDigit(rune(eid[4])) && unicode.IsDigit(rune(eid[5])) && unicode.IsDigit(rune(eid[6])) {
		return true
	}
	fmt.Printf("%s is not a valid ID.\n", eid)
	return false
}

func main() {
	var firstName, lastName, zipCode, employeeId string
	var errors int = 0

	fmt.Printf("Enter the first name: ")
	fmt.Scan(&firstName)

	fmt.Printf("Enter the last name: ")
	fmt.Scan(&lastName)

	fmt.Printf("Enter the ZIP code: ")
	fmt.Scan(&zipCode)

	fmt.Printf("Enter an employee ID: ")
	fmt.Scan(&employeeId)

	errors += btoi(!validateFirstName(firstName))
	errors += btoi(!validateLastName(lastName))
	errors += btoi(!validateZipCode(zipCode))
	errors += btoi(!validateEmployeeId(employeeId))
	
	if errors == 0 {
		fmt.Printf("There were no errors found.\n")
	}
}