// Write a function that simulates the behaviour of the Atoi function in Go. Atoi transforms a number defined as a string in a number defined as an int.
// Atoi returns 0 if the string is not considered as a valid number. For this exercise non-valid string chains will be tested. Some will contain non-digits characters.
// For this exercise the handling of the signs + or - does not have to be taken into account.
// This function will only have to return the int. For this exercise the error return of Atoi is not required.

package piscine

func BasicAtoi2(s string) int {
	o_number := 0
	c := 0
	checker := true
	a_s := []rune(s)
	for _, word := range a_s {
		if byte(word) >= 48 && byte(word) <= 57 {
			for i := '0'; i < word; i++ {
				c++
			}
			o_number = o_number*10 + c
			c = 0
		} else {
			checker = false
		}
	}
	if checker {
		return o_number
	} else {
		return 0
	}
}

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	fmt.Println(piscine.BasicAtoi2("12345"))
// 	fmt.Println(piscine.BasicAtoi2("0000000012345"))
// 	fmt.Println(piscine.BasicAtoi2("012 345"))
// 	fmt.Println(piscine.BasicAtoi2("Hello World!"))
// }

// OUTPUT

// $ go run .
// 12345
// 12345
// 0
// 0
// $
