// Write a function that returns the n rune of a string.
// In case of impossibilities, the function returns 0.

package piscine


func NRune(s string n int) rune {
	runes := []rune(s)
	len := 0
	for i := range runes{
		len = i
	}

	if n-1 >= 0 && n-1 <= len{
		return runes[n-1]
	}
	return 0
}