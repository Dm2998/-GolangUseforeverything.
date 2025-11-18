package main 

import "fmt"

var steps int = 21  //  string type variable declaration and initialization for number of steps

type Message string // custom type Message based on string

func (msg Message) Say() { fmt.Println(msg) } // method Say for type Message to print the message

func main () {
	msg := Message("Learning Go is fun!") // create a Message instance
	msg.Say() // call the Say method to print the message

	fmt.Println("Total Steps in Learning Go:", steps) // Print total steps
}