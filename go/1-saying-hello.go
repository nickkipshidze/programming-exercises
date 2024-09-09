package main

import "fmt"

func main() {
	var name string
	
	fmt.Printf("What is your name? ")
	fmt.Scan(&name)
	fmt.Printf("Hello, %s, nice to meet you!\n", name)
}