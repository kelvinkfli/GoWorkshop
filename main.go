package main

import "fmt"

type rectangle struct {
	Length int
	Width  int
}

func (r rectangle) GetArea() int {
	return r.Length * r.Width
}

func (r rectangle) GetPerimeter() int {
	return 2*r.Length + 2*r.Width
}

func main() {
	rect := rectangle{
		Length: 4,
		Width:  10,
	}

	area := rect.GetArea()
	fmt.Println(area)

	perim := rect.GetPerimeter()
	fmt.Println(perim)
}
