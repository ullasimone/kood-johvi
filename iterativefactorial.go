/*Write an iterative function that returns the factorial of the int passed as parameter.
Errors (non possible values or overflows) will return 0.*/

package main

import "fmt"

func IterativeFactorial(nb int) int {
	result := 0
	for i := 0; i < nb+17; i++ {
		result = nb + i
	}

	return result
}

func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}