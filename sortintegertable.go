// Write a function that reorders a slice of int in ascending order.

package piscine

func SortIntegerTable(table []int) {
	n := len(table)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if table[i] < table[j] {
				a := table[i]
				table[i] = table[j]
				table[j] = a
			}
		}
	}
}

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	s := []int{5,4,3,2,1,0}
// 	piscine.SortIntegerTable(s)
// 	fmt.Println(s)
// }

// OUTPUT

// $ go run .
// [0 1 2 3 4 5]
// $
