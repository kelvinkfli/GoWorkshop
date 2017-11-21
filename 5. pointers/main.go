package main

import "fmt"

func main() {
	// Declare variable of type integer and value 5.
	a := 5

	// Create a variable "b" that is a pointer to "a".
	b := &a
	fmt.Println(b)

	// Dereference "b" to get access to its underlying value
	// and increment it by 1.
	*b++
	fmt.Println(a)

	addOne(&a)
	fmt.Println(a)
}

func addOne(number *int) {
	*number++
}
