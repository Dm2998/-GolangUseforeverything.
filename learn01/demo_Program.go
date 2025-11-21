package main // Specifies the package name


import "math/rand" //  Imports the fmt package for formatted I/O

const MaxRnd = 100 // Constant declaration for maximum random number value

/*

Staat Random number then counts and retuns for the small and large numbers

*/

func StartRandomCount(n int) (int, int) {
	// Declare variables to hold counts
	var a, b int 
	// A for loop to control the number of iterations
	for i := 0; i < n; i++ {
		if rand.Intn(MaxRnd) < MaxRnd/2 {
			a = a + 1 // increment small number count 

		} else {

			b++ // increment large number count
		}
	}
	return a,b // return the counts, same as : b = b + 1
}


func main() {
	var num  = 1000 // Number of random numbers to generate
	x, y := StartRandomCount(num) // Call the function and get counts

	// Print the results
	print("Total Small Number: ", x, " + ", y, " = ", num, "\n") 
	println(x + y == num) // Verify the total count
}