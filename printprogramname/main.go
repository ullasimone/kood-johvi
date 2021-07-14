// Write a program that prints the name of the program.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []rune(os.Args[0])  // casting - type conversion; type translation; translate the string into an array of runes
	for _, i := range arg[0] { // no need of index, (skip it (for _,)) then translate the rune array into symbols
		z01.PrintRune(i)
	}
	z01.PrintRune(10)
}
