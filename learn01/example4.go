package main 

import "fmt"

var steps int = 30    // string type variable declaration and initialization for number of steps


type Message string 

    func (msg Message) String() string { return string("<<" + msg + ">>")
}
	func main() {

		msg := Message("Learning Go is my Plan for now!") // create a Message instance
		fmt.Println(msg, steps, "its A GOOD idea!") // Print message along with steps))
	}

// custom type Message based on string


