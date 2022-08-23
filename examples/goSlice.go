package main

import "fmt"

func main() {
	aSlice := []float64{}
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	fmt.Println()
	aSlice = append(aSlice, 1.1)
	aSlice = append(aSlice, 2.2)
	aSlice = append(aSlice, 3.3)
	fmt.Println(aSlice, len(aSlice), cap(aSlice))
	fmt.Println()
	t := make([]int, 4)
	t[0] = 1
	t[1] = 2
	t[2] = 3
	t[3] = 4
	fmt.Println(t, len(t), cap(t))
	fmt.Println()
	t = append(t, 5)
	fmt.Println(t, len(t), cap(t))
	fmt.Println()
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	for _, v := range twoD {
		for _, v2 := range v {
			fmt.Print(v2, " ")
		}
		fmt.Println()
	}
	fmt.Println()
	make2D := make([][]int, 2)
	fmt.Println(make2D)
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D)
}
