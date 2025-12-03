

package main

import "fmt"

func main() {
	months := make([]string, 12)
	months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	fmt.Println(months)

	q1 := months[0:3]
	q2 := months[3:6]

	fmt.Println("First quarter:", len(months), cap(months), months)
	fmt.Println("Second quarter:", len(q1), cap(q1), q2)

}