package main

import (
	"fmt"
	"strconv"
)

func main() {
	n := 10
	input1 := strconv.Itoa(n)
	fmt.Println(input1)
	fmt.Printf("type %T\n", input1)
	fmt.Println("############################")
	input2 := strconv.FormatInt(int64(n), 10)
	fmt.Println(input2)
	fmt.Printf("type %T\n", input2)
	fmt.Println("############################")
	input3 := string(n)
	fmt.Println(input3)
	fmt.Printf("type %T\n", input3)
}
