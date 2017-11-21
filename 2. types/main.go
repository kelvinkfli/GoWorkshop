package main

import "fmt"

func main() {
	// A declared variable of type 'string' without a value.
	var a string

	// Shorthand variable declaration with type inference.
	fruit := "pineapple"

	// Multiple variable declaration.
	number, boolean := 50, false

	fmt.Println(a, fruit, number, boolean)
}
