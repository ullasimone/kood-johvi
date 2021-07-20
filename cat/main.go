/*Write a program that behaves like a simplified cat command.
The options do not have to be handled.
If the program is called without arguments it should take the standard input (stdin) and print it back on the standard output (stdout).*/

package main

import (
	"io/ioutil"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	lens := 0
	argum := os.Args[1:]
	for range argum {
		lens++
	}
	for i := 1; i <= lens; i++ {
		file, err := ioutil.ReadFile(os.Args[i])
		strFile := string(file)
		if err != nil {
			string := "open " + os.Args[i] + ": no such file or directory"
			for _, v := range string {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		} else {
			for _, v := range strFile {
				z01.PrintRune(v)
			}
			z01.PrintRune('\n')
		}
	}
	if lens == 0 {
		for {
			text, _ := ioutil.ReadAll(os.Stdin)
			for _, v := range string(text) {
				z01.PrintRune(v)
			}
		}
	}
}
