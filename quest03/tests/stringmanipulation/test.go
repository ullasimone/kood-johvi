package main

import "fmt"

func main() {
	aString := "Hello"

	aStringChangeable := []byte(aString)

	aStringChangeable[0] = 'A'

	aStringFinalized := string(aStringChangeable)

	fmt.Println(aString)

	fmt.Println(aStringChangeable)

	fmt.Println(aStringFinalized)

}

//Ullas-MacBook-Air:Quest03 ullasimone$ go run test.go
//Hello
//[65 101 108 108 111]
//Aello
