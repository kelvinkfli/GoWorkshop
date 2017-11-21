package main

import (
	"fmt"

	"./array"
)

func main() {
	numbers := []int{1, 2, 3}
	sum, isOdd := getSum(numbers)
	fmt.Println(sum, isOdd)

	array.PrintNumbers(numbers)
}

func getSum(numbers []int) (int, bool) {
	var sum int
	var isOdd bool

	for _, n := range numbers {
		sum = sum + n
	}

	if sum%2 != 0 {
		isOdd = true
	}

	return sum, isOdd
}
