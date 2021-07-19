// Write a function that returns the n rune of a string.
// In case of impossibilities, the function returns 0.

package piscine

func NRune(s string, n int) rune {
	res := []rune(s)            // changes string into runes
	if n > 0 && n <= len(res) { // searching number is higher than 0 and lower than the length of the string
		return res[n-1] // return string with the index of number-1
	} else {
		return 0
	}
}
