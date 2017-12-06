package main

import (
	"fmt"
	"time"
)

func main() {
	// Lets run some asynchronous action in a loop!
}

func asyncAction(iteration int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Iteration: %d \n", iteration)
}
