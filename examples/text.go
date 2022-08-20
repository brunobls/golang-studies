package main

import "fmt"

func main() {
	aString := "Hello, World! €"
	fmt.Println("First character:", string(aString[0]))
	fmt.Println("#########################################")
	r := '€'
	fmt.Println("As an int32 value:", r)
	fmt.Printf("As a rune: %s and as a character: %c\n", r, r)
	fmt.Println("#########################################")
	for _, v := range aString {
		fmt.Printf("%x(%c) ", v, v)
	}
	fmt.Println()
	for _, v := range aString {
		fmt.Printf("%c", v)
	}
	fmt.Println()
}
