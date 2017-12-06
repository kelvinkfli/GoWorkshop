package main

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
	Base   int
	Height int
}

func (t triangle) area() int {
	return t.Base * t.Height / 2
}

func main() {
	s1 := square{
		Width:  2,
		Length: 2,
	}
	t1 := triangle{
		Base:   5,
		Height: 11,
	}

	// Lets write a function that compares the area of two shapes!
}
