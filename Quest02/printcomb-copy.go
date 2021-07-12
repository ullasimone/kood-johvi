/* Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.
These combinations are separated by a comma and a space.*/

package main

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	for a := '0'; a <= '7'; a++ {
		if a < b {
			for b := '1'; b <= '8'; b++ {
				if b < c {
					for c := '2'; c <= '9'; c++ {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						z01.PrintRune(',')
						z01.PrintRune(' ')

					}
				} else {
					c++
				}
			}
		} else {
			b++
			b := '1'
			c := '2'
		}
	}
}

func main() {
	PrintComb()
}

/*
package main

import "piscine"

func main() {
	piscine.PrintComb()
}

OUTPUT (incomplete)

$ go run . | cat -e
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789$
$*/
