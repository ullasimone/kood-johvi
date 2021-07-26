/*Write an iterative function that returns the factorial of the int passed as parameter.
Errors (non possible values or overflows) will return 0.*/

package piscine

func IterativeFactorial(nb int) int {
	result := 1
	// factorial number
	if nb == 1 || nb == 0 {
		return 1
		// i = 1, 'cause you can't get a valid number if you multiply with 0
	} else if nb > 1 && nb <= 20 {
		for i := 1; i < nb+1; i++ {
			result = result * i
		}
		// everything under 0
	} else {
		return 0
	}
	return result
}
