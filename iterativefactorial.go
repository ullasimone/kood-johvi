/*Write an iterative function that returns the factorial of the int passed as parameter.
Errors (non possible values or overflows) will return 0.*/

package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	} else {
		result := 1
		for i := 1; i < nb+1; i++ {
			result = nb * i
		}
		return result
	}
}

/*func main() {
	arg := 4
	fmt.Println(IterativeFactorial(arg))
}*/
