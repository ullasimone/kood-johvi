// Write a function that prints one by one the characters of a string on the screen.

package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, s := range s {
		z01.PrintRune(s)
	}
}

// package main

// import "piscine"

// func main() {
// 	piscine.PrintStr("Hello World!")
// }

// OUTPUT

// $ go run . | cat -e
// Hello World!%
// $
