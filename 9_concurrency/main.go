package main

import (
	"fmt"
	"time"
)

func main() {
	// Lets run some asynchronous action in a loop!
}

func asyncAction(number int) {
	time.Sleep(500 * time.Millisecond)
	fmt.Printf("Action: %d \n", number)
}
