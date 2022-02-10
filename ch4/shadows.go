package main

import "fmt"

func main() {
	// assigns value to x
	x := 10
	if x > 5 {
		fmt.Println(x)
		// shadows x with value 5 so long as in this block
		x := 5
		fmt.Println(x)
	}
	// exits block x value is now 10 again
	fmt.Println(x)
}
