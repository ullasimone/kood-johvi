// Write a function that will be formatted as below.
// This function will divide the int a and b.
// The result of this division will be stored in the int pointed by a.
// The remainder of this division will be stored in the int pointed by b.

package piscine

func UltimateDivMod(a *int, b *int) {
	tempA := *a / *b
	tempB := *a % *b
	*a = tempA
	*b = tempB
}

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 13
// 	b := 2
// 	piscine.UltimateDivMod(&a, &b)
// 	fmt.Println(a)
// 	fmt.Println(b)
// }

// OUTPUT

// $ go run .
// 6
// 1
// $
