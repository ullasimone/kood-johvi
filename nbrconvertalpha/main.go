/*Write a program that prints the corresponding letter in the n position of the latin alphabet, where n is each argument received.
For example 1 matches a, 2 matches b, etc. If n does not match a valid position of the alphabet or if the argument is not an integer, the program should print a space (" ").
A flag --upper should be implemented. When used, the program prints the result in upper case. The flag will always be the first argument.*/

package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args
	counter := 0
	index := 96
	for i, s := range arg {
		counter = i + 1
		if s == "--upper" {
			index = 64
		}
	}
	for j := 1; j <= counter-1; j++ {
		for _, l := range arg[j] {
			z01.PrintRune(l + rune(index))
		}
		z01.PrintRune(10)
	}
}
