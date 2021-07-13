// Write a function that counts only the letters of a string and that returns that count.
// White spaces or any other characters should not be counted.
// The letters are only the ones from the latin alphabet.

package piscine

func AlphaCount(s string) int {
	aString := []rune(s)
	counter := 0

	for i := range s {
		if aString[i] <= 'Z' && aString[i] >= 'A' || aString[i] <= 'z' && aString[i] >= 'a' {
			counter++
		}
	}
	return counter
}
