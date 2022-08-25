package main

import "fmt"

type aStructure struct {
	field1 complex128
	field2 int
}

func processPointer(x *float64) {
	*x = *x * *x
}

func returnPointer(x float64) *float64 {
	temp := 2 * x
	return &temp
}

func bothPointers(x *float64) *float64 {
	temp := 2 * *x
	return &temp
}

func main() {
	var f float64 = 2
	fmt.Println("Memory address of f:", &f)
	fmt.Printf("\n##############\n")

	// Pointer to f
	fP := &f
	fmt.Println("Memory address of f:", fP)
	fmt.Println("Value of f:", *fP)
	fmt.Printf("\n##############\n")

	// The value of f changes
	processPointer(fP)
	fmt.Printf("Value of f: %.2f\n", f)
	fmt.Printf("\n##############\n")

	// The value of f does not change
	x := returnPointer(f)
	fmt.Printf("Value of x: %.2f\n", *x)
	fmt.Printf("\n##############\n")

	// The value of f does not change
	xx := bothPointers(fP)
	fmt.Printf("Value of xx: %.2f\n", *xx)
	fmt.Printf("\n##############\n")

	// Check for empty structure
	var k *aStructure
	// This is nil because currently k points to nowhere
	fmt.Println(k)
	fmt.Printf("\n##############\n")

	// Therefore you are allowed to do this:
	if k == nil {
		k = new(aStructure)
	}
	fmt.Printf("\n##############\n")
	fmt.Printf("%+v\n", k)
	if k != nil {
		fmt.Println("k is not nil!")
	}
}
