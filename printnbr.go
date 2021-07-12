// Write a function that prints an int passed in parameter. All possible values of type int have to go through. You cannot convert to int64.

package piscine

import "github.com/01-edu/z01"

func PrintNbr(n int) {
	i := '0'
	if n == 0 {
		z01.PrintRune('0')
		return
	}
	for j := 1; j <= n%10; j++ {
		i++
	}
	for j := -1; j >= n%10; j-- {
		i++
	}
	if n/10 != 0 {
		PrintNbr(n / 10)
	}
	z01.PrintRune(i)
	return
}

func PrintNbr(n int) {
	if n < 0 {
		z01.PrintRune('-')
	}
	PrintNbr(n)
}

/*func main() {
	PrintNbr(-123)
	PrintNbr(0)
	PrintNbr(123)
	z01.PrintRune('\n')
}

OUTPUT

$ go run .
-1230123
$*/
