// Write a function that takes a pointer to a pointer to a pointer to an int as argument and gives to this int the value of 1.

package piscine

func UltimatePointOne(n ***int) {
	***n = 1
}

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 0
// 	b := &a
// 	n := &b
// 	piscine.UltimatePointOne(&n)
// 	fmt.Println(a)
// }

// OUTPUT

// $ go run .
// 1
// $
