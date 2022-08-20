package main

import "fmt"

func Print[T any](values []T) {
	for _, value := range values {
		fmt.Print(value, " ")
	}
	fmt.Println()
}

func main() {
	ints := []int{1, 2, 3, 4, 5}
	Print(ints)
	strings := []string{"a", "b", "c", "d", "e"}
	Print(strings)
}
