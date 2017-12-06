package main

import "fmt"

func main() {
	// Declare a channel.
	ch := make(chan string)

	// Send a message on that channel!
	go sendMessage("HELLO!", ch)

	// Receive a message on that channel!
	msg := <-ch
	fmt.Println(msg)
}

func sendMessage(msg string, ch chan string) {
	ch <- msg
}
