package main

import (
	"os"

	"github.com/01-edu/z01"
)

func printStr(s string) {
	arrayS := []rune(s)
	counter := 0
	for r := range arrayS {
		counter = r + 1
	}
	for r := 0; r < counter; r++ {
		z01.PrintRune(arrayS[r])
	}
	z01.PrintRune('\n')
}

func isEven(nbr int) bool {
	if nbr%2 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	args := os.Args[1:]
	count := 0
	for i := range args {
		count = i + 1
	}
	if isEven(count) {
		printStr("I have an even number of arguments")
	} else {
		printStr("I have an odd number of arguments")
	}
}
