package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	d := time.Now().UnixNano()
	rand.Seed(d)
	fmt.Printf("\nd: %d\n", d)
}

func main() {
	for i := 0; i < 100; i++ {
		fmt.Printf("%d", rand.Intn(20))
	}
	fmt.Println()
}
