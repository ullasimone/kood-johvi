// Write a function that counts the runes of a string and that returns that count.

package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}