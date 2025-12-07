package main 

import "fmt"


var (
	empty struct{}
	car struct { make, model string }
	currency struct {
		name, country string 
		code          int 

	}

	node struct {
		edges []string
		weight int
	}

	person struct {
		name string
		address struct {
			street       string
			city, state  string
			postal       string
		}
	}
)

func main() {
	fmt.Println("E,pty Struct:", empty)
	fmt.Println("Car Struct:", car)
	fmt.Println("Currentcy Strunt:", currency)
	fmt.Println("Node Struct:", node)
	fmt.Println("Person Struct:", person) 
}