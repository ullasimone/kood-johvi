/*Write a program that behaves like a simplified cat command.
The options do not have to be handled.
If the program is called without arguments it should take the standard input (stdin) and print it back on the standard output (stdout).*/

package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 1 {
		return
	} else {
		os.Args = os.Args[1:]

		for _, v := range os.Args {
			file, err := os.Open(v)
			if err != nil {
				fmt.Println(err)
				return
			} else {
				data := make([]byte, 443)
				file.Read(data)
				if len(os.Args) == 1 {
					fmt.Println(string(data))
				} else {
					fmt.Println(string(data))
				}
				file.Close()
			}
		}
	}
}
