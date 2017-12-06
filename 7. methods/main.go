package main

import "fmt"

type square struct {
	Width  int
	Length int
}

func (s square) area() int {
	return s.Width * s.Length
}

func (s square) isLarge() bool {
	if s.Width > 5 && s.Length > 5 {
		return true
	}
	return false
}

func main() {
	s := square{
		Width:  6,
		Length: 6,
	}

	// Lets call the "area" method on our square.
	area := s.area()
	fmt.Println(area)

	// Lets call the "isLarge" method on our square.
	isLarge := s.isLarge()
	fmt.Println(isLarge)
}
