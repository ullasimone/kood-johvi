// Write a function that returns true if the string passed in parameter only contains uppercase characters, and that returns false otherwise.

package piscine

func IsUpper(s string) bool {
	runeArray := []rune(s)
	runeCounter := arrayCount(s)
	counter := 0
	for _, i := range runeArray {
		// loop only capital letters
		if i <= 'Z' && i >= 'A' {
			counter++
		}
	}
	if counter == runeCounter {
		return true
	}
	return false
}
