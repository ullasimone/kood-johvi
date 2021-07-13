// Write a function that returns the n rune of a string.
// In case of impossibilities, the function returns 0.

package piscine

func NRune(s string, n int) rune {
	sRune := []rune(s)
	return sRune[n-1]
}
