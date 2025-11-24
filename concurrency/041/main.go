// Exercise: Channels - Closing

// Create a string channel "c" (make it a buffered channel)
// Add 2 different strings directly into that channel.
// Close the channel with the close() statement and read a quote from the channel, Can you read it?

package main

import "fmt"

func main() {
	c := make(chan string)

	go func() {
        c <- "string 1"
	    c <- "string 2"
			close(c)
	}()

	msg1, ok := <-c

	if ok {
		fmt.Println(msg1)
	}



	msg2, yes := <-c

	if yes {
		fmt.Println(msg2)
	}
}
