// Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number defined as a string in a number defined as an int.
// Atoi returns 0 if the string is not considered as a valid number. For this exercise only valid string will be tested. They will only contain one or several digits as characters.
// For this exercise the handling of the signs + or - does not have to be taken into account.
// This function will only have to return the int. For this exercise the error return of Atoi is not required.

package piscine

func BasicAtoi(s string) int {
	aNumber := 0
	a := 0
	s := []rune(s)
	for _, word := range s {
		for i := '0'; i < word; i++ {
			a++
		}
		aNumber = aNumber*10 + a
		a = 0
	}
	return aNumber
}
