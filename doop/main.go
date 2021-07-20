/*Write a program that is called doop.
The program has to be used with three arguments:
A value
An operator, one of : +, -, /, *, %
Another value
In case of an invalid operator, value, number of arguments or an overflow, the programs prints nothing.
The program has to handle the modulo and division operations by 0 as shown on the output examples below.*/

package main

import (
	"fmt"
	"os"
	"strconv"
)

func isOperator(s string, a []string) bool {
	for _, val := range a {
		if s == val {
			return true
		}
	}
	return false
}

func main() {
	op := []string{"+", "*", "-", "/", "%"}
	args := os.Args[1:]
	if len(args) != 3 {
	} else {
		if isOperator(args[1], op) {
			n1, err1 := strconv.Atoi(args[0])
			n2, err2 := strconv.Atoi(args[2])
			if err1 == nil && err2 == nil {
				switch args[1] {
				case "+":
					fmt.Println(n1 + n2)
				case "-":
					fmt.Println(n1 - n2)
				case "*":
					fmt.Println(n1 * n2)
				case "/":
					if n2 == 0 {
						fmt.Println("Error: no division by 0")
					} else {
						fmt.Println(n1 / n2)
					}
				case "%":
					if n2 == 0 {
						fmt.Println("Error: no modulo by 0")
					} else {
						fmt.Println(n1 % n2)
					}

				}
			} else {
				fmt.Println("1")
			}

		} else {
			fmt.Println("0")
		}
	}
}
