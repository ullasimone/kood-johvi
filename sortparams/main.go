// Write a program that prints the arguments received in the command line in ASCII order.
// OUTPUT go run sortparams/main.go 1 a 2 A 3 b 4 C

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	counter := 0
	for s := range arg {
		counter = s + 1
	}
	for i := 1; i < counter; i++ {
		for j := i + 1; j < counter; j++ {
			if arg[i] > arg[j] {
				arg[i], arg[j] = arg[j], arg[i]
			}
		}
	}
	for j := 1; j <= counter-1; j++ {
		for _, l := range arg[j] {
			z01.PrintRune(l)
		}
		z01.PrintRune(10)
	}
}
