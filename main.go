package main

import "fmt"

func main() {
	// A declared variable of type 'string' without a value.
	var a string

	// A declared variable that has inferred its type from the value.
	var animal = "cat"

	// Shorthand variable declaration.
	fruit := "carrot"

	// Multiple variable declaration.
	number, boolean := 50, false

	var (
		x = 1
		y = 2
		z = "3"
	)

	fmt.Println(a, animal, fruit, number, boolean, x, y, z)
}

/*
- Statically and strongly typed.
*/
