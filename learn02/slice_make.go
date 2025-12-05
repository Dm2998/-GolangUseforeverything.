

package main

import "fmt"

func main() {
	months := make([]string, 12)
	months = []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	fmt.Println(months)

	q1 := months[0:3]
	q2 := months[3:6]
	q3 := months[6:9]
	q4 := months[9:12]

	summery := months[2:4:10]

	fmt.Println("First quarter:", len(months), cap(months), months)
	fmt.Println("Second quarter:", len(q1), cap(q1), q2)
	fmt.Println("Third quarter:", len(q2), cap(q2), q3)
	fmt.Println("Fourth quarter:", len(q3), cap(q3), q4)
	fmt.Println("Summery:", len(summery), cap(summery), summery)

}