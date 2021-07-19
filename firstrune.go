// Write a function that returns the first rune of a string.

package piscine

func FirstRune(s string) rune {
	fR := []rune(s) // changes string into runes
	return fR[0]    // returns strings first rune; its first rune is in the place "0"
}
