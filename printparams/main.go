// Write a program that prints the arguments received in the command line.
// OUTPUT go run printparams/main.go choumi is the best cat

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	// finds arg char count and strings themselves
	for i, s := range arg {
		// i cannot be 0 'cause 0 == programs name
		if i != 0 {
			// don't need to know the index, matters the found strings rune symbols
			for _, r := range s {
				// print symbols
				z01.PrintRune(r)
			} // give new live to every arg
			z01.PrintRune('\n')
		}
	}
}
