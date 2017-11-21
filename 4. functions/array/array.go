package array

import "fmt"

// PrintNumbers iteratively prints each number in the provided slice.
func PrintNumbers(numbers []int) {
	for _, n := range numbers {
		fmt.Println(n)
	}
}
