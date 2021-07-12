package piscine

func RecursivePower(nb int, power int) int {
	if power > 0 {
		return nb * RecursivePower(nb, power-1)
	} else {
		return 1
	}
	return 1
}

/*func main() {
	fmt.Println(RecursivePower(4, 3))
}*/
