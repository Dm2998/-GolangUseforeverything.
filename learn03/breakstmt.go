

package main

import "fmt"

var words = [][]string{
	{"break", "lake", "make", "go", "take", "strong", "bring"},
}

func search(word string) {
DoSearch:
	for i := 0; i < len(words); i++ {
		for k := 0; k < len(words[i]); k++ {
			if words[i][k] == word {
				fmt.Println("Found:", word)
				break DoSearch
			}
		}
	}
}

func main() {
	search("go")
}