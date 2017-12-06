package main

import (
	"fmt"
)

func main() {
	// Slices.
	var animals []string
	animals = []string{"gophers", "cats", "dogs"}

	// Let append some animals to this slice.
	animals = append(animals, "llamas")

	// Loop over the slice we just made.
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

	// Lets add a new key/value pair to the map.
	fruits["pineapples"] = 99

	// Loop over the map we just made.
	for key, value := range fruits {
		fmt.Println(key, value)
	}
}
