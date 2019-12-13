package main

import "fmt"

func main() {
	// prepare two variables d as how many times dividable by 100 n is the nth smallest int
	var d, n int
	// scan inputs
	fmt.Scan(&d, &n)
	// switch process by d = 0, 1, 2
	switch d {
	// when individable by 100
	case 0:
		// if n = 100 we need to skip 100 because it is dividable
		if n == 100 {
			fmt.Printf("%d", n+1)
		} else {
			fmt.Printf("%d", n)
		}
	// when dividable by 100 for once
	case 1:
		// if n = 100 we need to skip 100 because it is twice dividable
		if n == 100 {
			fmt.Printf("%d", (n+1)*100)
		} else {
			fmt.Printf("%d", n*100)
		}
	// when dividable by 100 for twice
	case 2:
		// if n = 100 we need to skip 100 because it is three times dividable
		if n == 100 {
			fmt.Printf("%d", (n+1)*10000)
		} else {
			fmt.Printf("%d", n*10000)
		}
	}
}
