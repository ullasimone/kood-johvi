// Write a program that prints the arguments received in the command line.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	// finds arg char count and chars themselves (symbols)
	for i, s := range arg {
		// i cannot be 0 'cause 0 == programs name
		if i != 0 {
			// don't need to know the index, matters the string rune symbols
			for _, r := range s {
				// print symbols
				z01.PrintRune(r)
			} // give new live to every arg
			z01.PrintRune('\n')
		}
	}
}
