// Write a function that returns true if the string passed in parameter only contains alphanumerical characters or is empty, and returns false otherwise.

package piscine

func IsAlpha(s string) bool {
	runeArray := []rune(s)
	runeCounter := arrayCount(s)
	counter := 0
	for _, i := range runeArray {
		if i >= 'a' && i <= 'z' || i >= 'A' && i <= 'Z' || i >= '0' && i <= '9' {
			counter++
		}
	}
	if counter == runeCounter {
		return true
	}
	return false
}
