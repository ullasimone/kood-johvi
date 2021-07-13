// Write a function that returns true if the string passed as a parameter only contains numerical characters, and returns false otherwise.

package piscine

func IsNumeric(s string) bool {
	runeArray := []rune(s)
	runeCounter := arrayCount(s)
	counter := 0
	for _, i := range runeArray {
		if i >= '0' && i <= '9' {
			counter++
		}
		if counter == runeCounter {
			return true
		}

	}
	return false
}
