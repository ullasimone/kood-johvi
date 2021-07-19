package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	missing := "File name missing"
	tooMany := "Too many arguments"
	file := "quest8.txt"
	if len(os.Args) < 2 {
		fmt.Println(missing)
		return
	}
	if len(os.Args) > 2 {
		fmt.Println(tooMany)
		return
	}
	data, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(data))
}
