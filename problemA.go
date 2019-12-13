package main

import "fmt"

func main() {
	// prepare two variables
	var a, b int
	// scan inputs
	fmt.Scan(&a, &b)
	// check if one of the inputs is bigger than 8
	if a < 9 && b < 9 == true {
		// both inputs are smaller or equal to 8, print Yay!
		fmt.Println("Yay!")
	} else {
		// atleast one of the inputs are bigger than 8, print :(
		fmt.Println(":(")
	}
}
