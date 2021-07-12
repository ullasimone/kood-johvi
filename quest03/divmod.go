// Write a function that will be formatted as below.
// This function will divide the int a and b.
// The result of this division will be stored in the int pointed by div.
// The remainder of this division will be stored in the int pointed by mod.

package piscine

func DivMod(a int, b int, div *int, mod *int) {
	*div = a / b
	*mod = a % b
}

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	a := 13
// 	b := 2
// 	var div int
// 	var mod int
// 	piscine.DivMod(a, b, &div, &mod)
// 	fmt.Println(div)
// 	fmt.Println(mod)
// }

// OUTPUT

// $ go run .
// 6
// 1
// $
