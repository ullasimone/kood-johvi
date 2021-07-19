package piscine

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, s := range s { // loop searches every character in string
		z01.PrintRune(s) // prints string
	}
}
