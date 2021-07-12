// Write a function that returns the square root of the int passed as parameter, if that square root is a whole number. Otherwise it returns 0.

package piscine

func Sqrt(nb int) int {
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

/*func main() {
	fmt.Println(Sqrt(4))
	fmt.Println(Sqrt(3))
}*/
