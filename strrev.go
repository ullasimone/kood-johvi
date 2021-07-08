// Write a function that reverses a string.
// This function will return the reversed string.

package piscine

func StrRev(s string) string {
	reverse := ""
	for _, i := range s {
		reverse = string(i) + reverse
	}
	return reverse
}
