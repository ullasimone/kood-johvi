// Write a function that counts the runes of a string and that returns that count.

package piscine

func StrLen(s string) int {
	i := 0
	for range s {
		i++
	}
	return i
}

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	l := piscine.StrLen("Hello World!")
// 	fmt.Println(l)
// }

// OUTPUT

// $ go run .
// 12
// $
