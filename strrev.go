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

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := "Hello World!"
// 	s = piscine.StrRev(s)
// 	fmt.Println(s)
// }

// OUTPUT

// $ go run .
// !dlroW olleH
// $
