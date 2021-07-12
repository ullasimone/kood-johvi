package main

import "github.com/01-edu/z01"

// single rune

// func main() {
// 	var aRune rune = 'A'

// 	z01.PrintRune(aRune)
// } //prints A

// character from a string

// func main() {
// 	var aRune string = "Hello"

// 	z01.PrintRune(rune(aRune[0]))
// } // prints H

// print whole string

func main() {
	var aRune string = "Hello"

	z01.PrintRune(rune(aRune[0]))
	z01.PrintRune(rune(aRune[1]))
	z01.PrintRune(rune(aRune[2]))
	z01.PrintRune(rune(aRune[3]))
	z01.PrintRune(rune(aRune[4]))
	z01.PrintRune('\n')
} // prints Hello
