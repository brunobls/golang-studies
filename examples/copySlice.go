package main

import "fmt"

func main() {
	a1 := []int{1}
	a2 := []int{-1, -2}
	a5 := []int{10, 11, 12, 13, 14}
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5)

	fmt.Printf("\n\n")

	// copy(destination, input)
	// len(a2) > len(a1)
	copy(a1, a2)
	fmt.Println("a1", a1)
	fmt.Println("a2", a2)

	fmt.Printf("\n\n")

	// len(a5) > len(a1)
	copy(a1, a5)
	fmt.Println("a1", a1)
	fmt.Println("a5", a5)

	fmt.Printf("\n\n")

	// len(a2) < len(a5) -> OK
	copy(a5, a2)
	fmt.Println("a2", a2)
	fmt.Println("a5", a5)

	fmt.Printf("\n\n")
}

// Go oferece a função copy() para copiar um array
// existente para um slice ou um slice existente para
// outro slice. No entanto, o uso de copy() pode ser
// complicado porque a fatia de destino não é expandida
// automaticamente se a fatia de origem for maior que a
// fatia de destino. Além disso, se a fatia de destino for
// maior que a fatia de origem, copy() não esvaziará os
// elementos da fatia de destino que não foram copiados.
