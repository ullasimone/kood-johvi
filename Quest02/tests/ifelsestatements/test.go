package main

import "fmt"

// "In most countries you are tall.""

// func main() {
// 	height := 185
// 	if height >= 180 {
// 		fmt.Println("In most countries you are tall.")
// 	}
// }

// Using else "You're of average height."
// func main() {
// 	height := 175
// 	if height >= 180 {
// 		fmt.Println("In most countries you are tall.")
// 	else {
// 		fmt.Println("You're of average height.")
// 	}
// }

// Using extra if else "You are very tall."
func main() {
	height := 210
	if height >= 180 && height < 200 {
		fmt.Println("In most countries you are tall.")
	} else if height >= 200 {
		fmt.Println("You are very tall.")
	} else {
		fmt.Println("You're of average height.")
	}
}
