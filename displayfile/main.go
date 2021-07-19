// Write a program that displays, on the standard output, the content of a file given as argument.

package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	file := "quest8.txt"

	if len(os.Args) < 2 {
		fmt.Println("File name missing")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments")
		return
	}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(data))
}
