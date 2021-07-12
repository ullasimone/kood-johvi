package main

import "fmt"

// initializer inside of "for"

// func main() {
// 	for i := 0; i <= 10; i++//incrementation {
// 		fmt.Println(i)
// 	}
// }

// initializer outside of "for"

func main() {
	i := 0

	for i <= 10 {
		fmt.Println(i)
		i++
	}
}
