// Learning Go: Chapter 2, math.go


package main

import (  
	"math"
	"fmt"
)

// Struct
type triangle struct {
	name 	 string
	a, b, c float64
}

func (t *triangle) area() float64  {
	return 0.5 * t.a * t.b
}

func 

// Function 

func main() {
	t := &triangle{"Right", 1, 2, 3}
	fmt.Println(t.shapeInfo(t))

}

