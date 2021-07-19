// Write a program that prints the arguments received in the command line in reverse order.
// go run revparams/main.go choumi is the best cat

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
	for j := counter - 1; j >= 1; j-- {
		for _, l := range arg[j] {
			z01.PrintRune(l)
		}
		z01.PrintRune(10)
	}
}
