package main

import (
	"fmt"
	"strings"
)

func isAnagram(first, second string) (bool) {
	if len(first) != len(second) {
		return false
	}

	for i := 0; i < len(first); i++ {
		if !strings.Contains(second, string(first[i])) {
			return false
		}
	}

	return true
}

func main() {
	var first, second string

	fmt.Println("Enter two strings and I'll tell you if they are anagrams:")

	fmt.Printf("Enter the first string: ")
	fmt.Scan(&first)

	fmt.Printf("Enter the second string: ")
	fmt.Scan(&second)

	anagram := isAnagram(first, second)

	if anagram == true {
		fmt.Printf("\"%s\" and \"%s\" are anagrams.\n", first, second)
	} else {
		fmt.Printf("Not anagrams.\n")
	}
}