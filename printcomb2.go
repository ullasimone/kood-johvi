package piscine

import "github.com/01-edu/z01"

func PrintComb2() {
	var a rune = '0'
	var b rune = '0'
	var c rune = '0'
	var d rune = '0'

	for a = '0'; a <= '9'; a = a + 1 {
		for b = '0'; b <= '9'; b = b + 1 {
			for c = '0'; b <= '9'; c = c + 1 {
				for d = '0'; d <= '9'; d = d + 1 {
					z01.PrintRune(a)
					z01.PrintRune(b)
					z01.PrintRune(' ')
					z01.PrintRune(c)
					z01.PrintRune(d)
					z01.PrintRune(',')
					if a < '9' || b < '8' || c == '9' || d == '9' {
						z01.PrintRune(',')
						z01.PrintRune(' ')
					}
				}
			}
		}
	}
	z01.PrintRune('\n')
}
