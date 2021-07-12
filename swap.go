// Write a function that takes two pointers to an int (*int) and swaps their contents.

package piscine

func Swap(a *int, b *int) {
	tempA := *a
	tempB := *b
	*a = tempB
	*b = tempA
}

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 0
// 	b := 1
// 	piscine.Swap(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// OUTPUT

// $ go run .
// 1
// 0
// $
