package main

import (
	"fmt"
)

func main() {
	// Slices (arrays).
	var animals []string
	animals = []string{"gophers", "cats", "dogs"}

	animals = append(animals, "llamas")

	for i := 0; i < len(animals); i++ {
		a := animals[i]
		fmt.Println(a)
	}

	// Maps.
	var fruits map[string]int
	fruits = map[string]int{
		"apples":  1,
		"oranges": 2,
		"pears":   3,
	}

	fruits["pineapples"] = 99

	for key, value := range fruits {
		fmt.Println(key, value)
	}
}
