// Write a program that displays, on the standard output, the content of a file given as argument.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	args := os.Args[1:]
	if 0 == len(args) {
		fmt.Println("File name missing")
	} else if len(args) >= 2 {
		fmt.Println("Too many arguments")
	} else {
		file, err := ioutil.ReadFile(args[0])
		if err != nil {
			fmt.Println(err)
		}
		fmt.Print(string(file))
	}
}
