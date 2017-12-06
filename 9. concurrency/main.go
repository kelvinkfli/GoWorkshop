package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 10; i++ {
		go asyncAction(i)
	}

	time.Sleep(5 * time.Second)
}

func asyncAction(iteration int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Iteration: %d \n", iteration)
}
