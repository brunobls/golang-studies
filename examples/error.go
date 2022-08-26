package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	err := check(0, 10)
	if err == nil {
		fmt.Println("No error")
	} else {
		fmt.Println(err)
	}

	err = check(0, 0)
	if err.Error() == "a and b cannot be zero" {
		fmt.Println("Custom error detected!")
		fmt.Println(err.Error())
	} else {
		fmt.Println(err)
	}

	err = formattedError(0, 0)
	if err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("-123")
	if err == nil {
		fmt.Println("Int value is", i)
	}

	i, err = strconv.Atoi("Y123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Int value is", i)
	}

}

func check(a, b int) error {
	if a == 0 && b == 0 {
		return errors.New("a and b cannot be zero")
	}

	return nil
}

func formattedError(a, b int) error {
	if a == 0 && b == 0 {
		return fmt.Errorf("a %d and b %d cannot be zero. UserID: %d", a, b, os.Getuid())
	}

	return nil
}
