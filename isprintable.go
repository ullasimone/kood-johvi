// Write a function that returns true if the string passed as a parameter only contains printable characters, and returns false otherwise.

package piscine

func IsPrintable(s string) bool {
	runeArray := []rune(s)
	for _, i := range runeArray {
		if i >= 0 && i <= 31 {
			return false
		}
	}
	return true
}
