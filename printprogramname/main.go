// Write a program that prints the name of the program.
// casting - type conversion; type translation; translate the string into an array of runes
// no need of index, (skip it (for _,)) then translate the rune array into symbols

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := []rune(os.Args[0])
	for i := 2; i <= len(arg)-1; i++ {
		z01.PrintRune(rune(arg[i]))
	}
	z01.PrintRune(10)
}
