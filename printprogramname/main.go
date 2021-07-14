// Write a program that prints the name of the program.

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	for _, res := range arg[0] {
		z01.PrintRune(res)
	}
	z01.PrintRune('\n')
}
