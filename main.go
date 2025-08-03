package main

import "fmt"

func main() {
	x := 10
	p := &x // p is a pointer to x
	fmt.Println("Value of x:", x)
	fmt.Println("Address of x:", &x)
	fmt.Println("Value of p (address of x):", p)
	fmt.Println("Value at address p (*p):", *p)
	*p = 20 // Change value of x through pointer
	fmt.Println("New value of x:", x)
}
