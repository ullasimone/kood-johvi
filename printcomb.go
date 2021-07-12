/* Write a function that prints, in ascending order and on a single line: all unique combinations of three different digits so that, the first digit is lower than the second, and the second is lower than the third.
These combinations are separated by a comma and a space.*/

package main

import (
	"github.com/01-edu/z01"
)

func PrintComb() {
	var (
		a rune = '0'
		b rune = '1'
		c rune = '2'
	)
	for a = '0'; a <= '7'; a++ {
		for b = '1'; b <= '8'; b++ {
			for c = '2'; c <= '9'; c++ {
				if a < b {
					if b < c {
						z01.PrintRune(a)
						z01.PrintRune(b)
						z01.PrintRune(c)
						if a == '7' && b == '8' && c == '9' {
							z01.PrintRune('\n')
						} else {
							z01.PrintRune(',')
							z01.PrintRune(' ')
						}
					}
				}
			}
		}
	}
}

func main() {
	PrintComb()
}

/*OUTPUT (incomplete)

$ go run . | cat -e
012, 013, 014, 015, 016, 017, 018, 019, 023, ..., 689, 789$
$*/
