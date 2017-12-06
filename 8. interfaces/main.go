package main

import "fmt"

// Interfaces are another type, just like integers, strings, or defined structs. Instead of
// having defining typed fields, we define method signatures.

type square struct {
	Width  int
	Length int
}

func (s square) area() int {
	return s.Length * s.Width
}

type triangle struct {
	Width  int
	Height int
}

func (t triangle) area() int {
	return t.Width * t.Height / 2
}

type shape interface {
	area() int
}

func main() {
	s1 := square{
		Width:  2,
		Length: 2,
	}
	t1 := triangle{
		Width:  5,
		Height: 11,
	}

	largerShape := hasLargerArea(s1, t1)
	fmt.Println(largerShape)
}

func hasLargerArea(s1 shape, s2 shape) shape {
	if s1.area() > s2.area() {
		return s1
	}
	return s2
}
