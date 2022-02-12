package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// calling simple function
	results := div(5, 2)
	fmt.Println(results)

	// calling variadic function
	fmt.Println(addToBase(3))
	fmt.Println(addToBase(3, 2))
	fmt.Println(addToBase(3, 2, 4, 5, 6, 8))
	a := []int{4, 3}
	fmt.Println(addToBase(3, a...))
	fmt.Println(addToBase(3, []int{1, 2, 3, 4, 5}...))

	// calling function that returns multiple values
	result, remainder, err := divAndRemainder(5, 2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(result, remainder)
}

// simple function declaration if params are same tyep can be declared after last param
func div(num int, denom int) int {
	if denom == 0 {
		return 0
	}
	return num / denom
}

// showing variadic param
func addToBase(base int, vals ...int) []int {
	out := make([]int, 0, len(vals))
	for _, v := range vals {
		out = append(out, base+v)
	}
	return out
}

// simple function declaration if params are same tyep can be declared after last param
func divAndRemainder(num int, denom int) (int, int, error) {
	if denom == 0 {
		return 0, 0, errors.New("Cannot devide by zero")
	}
	return num / denom, num % denom, nil
}
