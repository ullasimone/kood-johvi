package main

import "github.com/01-edu/z01"


func PrintComb() {
	for i := 0; i <= 7; i++ {
		for j := i+1; j <= 8; j++ {
			for k := j+1; k <= 9; k++ {
				if i==7 && j==8 && k==9 {
					z01.PrintRune("%d%d%d\n",i,j,k)
				}else{
					z01.PrintRune("%d%d%d, ",i,j,k)
					}
				}
			}
		}
	}

}
