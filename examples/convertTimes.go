package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()

	// Local time
	loc, _ := time.LoadLocation("Local")
	fmt.Printf("Current Location: %s\n", now.In(loc))

	fmt.Println()

	// NY
	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	fmt.Println()

	// London
	loc, _ = time.LoadLocation("Europe/London")
	fmt.Printf("London Time: %s\n", now.In(loc))

	fmt.Println()

	// Tokyo
	loc, _ = time.LoadLocation("Asia/Tokyo")
	fmt.Printf("Tokyo Time: %s\n", now.In(loc))

	fmt.Println()

	// Shanghai
	loc, _ = time.LoadLocation("Asia/Shanghai")
	fmt.Printf("Shanghai Time: %s\n", now.In(loc))

	fmt.Println()
}
