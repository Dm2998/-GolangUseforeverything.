package main 

import "fmt"

var dm = make (chan int) // Declare a channel for integer communication

func Pings() {
	for i := 0; i < 5; i ++ {
		fmt.Println("Ping", i)

	}
	dm <- 1 // Send a signal to the channel
}

func main() {
	go Pings() // Start the Pings function as a goroutine
	fmt.Println("Ping")
	<- dm // Wait for the signal from the channel
	fmt.Println("Pong") 
}