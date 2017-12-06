package main

import (
	"fmt"

	"./array"
)

func main() {
	numbers := []int{1, 2, 3}
	sum := getSum(numbers)
	fmt.Println(sum)

	array.PrintNumbers(numbers)
}

// Lets declare a function that returns the sum of all numbers passed in
// through a slice.
func getSum(numbers []int) int {
	var sum int

	for _, n := range numbers {
		sum += n
	}

	return sum
}
