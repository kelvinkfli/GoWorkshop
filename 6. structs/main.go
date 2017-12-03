package main

import "fmt"

// Sometimes, primitive types aren't enough to communicate what we want to do with a function or application.
// Go's answer to this problem is the struct.  It can be seen as a custom type that groups other typed fields
// together.
type square struct {
	Width  int
	Length int
}

func (s square) area() int {
	return s.Length * s.Width
}

func main() {
	var s square

	s = square{
		Width:  5,
		Length: 5,
	}

	// Lets print out these fields!
	fmt.Printf("width: %d \n", s.Width)
	fmt.Printf("length: %d \n", s.Length)

	// Let call the "area" method on our square.
	area := s.area()
	fmt.Printf("area: %d \n", area)
}
