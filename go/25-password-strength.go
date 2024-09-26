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

func passwordValidator(password string) (int) {
	var numbers, letters, symbols bool = false, false, false

	for _, char := range password {
		if unicode.IsDigit(char) {
			numbers = true
		}
		if unicode.IsLetter(char) {
			letters = true
		}
		if unicode.IsSymbol(char) || unicode.IsPunct(char) {
			symbols = true
		}
	}

	return btoi(numbers) + btoi(letters)*2 + btoi(symbols)
}

func passwordStatus(password string, score int) {
	switch score {
		case 1:
			fmt.Printf("The password \"%s\" is a very weak password.\n", password)
		case 2:
			fmt.Printf("The password \"%s\" is a weak password.\n", password)
		case 3:
			fmt.Printf("The password \"%s\" is a strong password.\n", password)
		case 4:
			fmt.Printf("The password \"%s\" is a very strong password.\n", password)
	}
}

func main() {
	passwordStatus("12345", passwordValidator("12345"))
	passwordStatus("abcdef", passwordValidator("abcdef"))
	passwordStatus("abc123xyz", passwordValidator("abc123xyz"))
	passwordStatus("1337h@xor!", passwordValidator("1337h@xor!"))
}
