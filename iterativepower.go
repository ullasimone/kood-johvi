package piscine

func IterativePower(nb int, power int) int {
	if power != 0 {
		return (nb * IterativePower(nb, power-1))
	} else {
		return 1
	}
}

/*func main() {
	fmt.Println(IterativePower(4, 3))
}*/
