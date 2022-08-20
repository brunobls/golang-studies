package main

import "fmt"

func main() {
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))
	fmt.Println()
	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))
	fmt.Println()
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice)
	fmt.Println()
	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)
	fmt.Println()
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
	fmt.Println()
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	fmt.Println(aSlice)
	fmt.Println()
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice))
}
