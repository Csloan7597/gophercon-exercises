package main

import "fmt"

func process(messages chan string, quit chan struct{}) {
	// Loop through your messages
	for m := range messages {
		// print the message with a for loop using range
		fmt.Println(m)
	}

	close(quit)
}

func main() {
	// declare the messages channel of type string and capacity of 5
	messages := make(chan string, 5)

	// declare a signal channel
	done := make(chan struct{}, 1)

	// launch process in a goroutine
	go process(messages, done)

	// declare 5 fruits in a []string
	fruits := []string{"apple", "banana", "cherry", "kiwi", "mango"}

	// loop through the fruits and send them to the messages channel
	for _, f := range fruits {
		messages <- f
	}

	// close the messages channel
	close(messages)

	// wait for everything to finish
	_ = <- done

	fmt.Println("done")
}