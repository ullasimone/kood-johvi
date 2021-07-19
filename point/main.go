/*Create a .go file.
The code below has to be copied in that file.
The necessary changes have to be applied so that the program works.
The program must be submitted inside a folder with the name point.*/

package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}

	setPoint(points)

	z := "x=42, y=21"

	for _, i := range z {
		z01.PrintRune(i)
	}

	z01.PrintRune('\n')
}
