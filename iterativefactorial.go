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
		return result
		// everything under 0
	} else {
		return 0
	}
}

/*func main() {
	result1 := IterativeFactorial(4)
	fmt.Println(result1)
}*/
