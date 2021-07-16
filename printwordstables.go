/*Write a function that prints the words of a string slice that will be returned by a function SplitWhiteSpaces.
Each word is on a single line (each word ends with a \n).*/

package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(a []string) {
	for _, s := range a {
		for _, char := range []rune(s) {
			z01.PrintRune(char)
		}
		z01.PrintRune('\n')
	}
}
