package piscine

func IterativePower(nb int, power int) int {
	result := nb
	for i := 1; i < power; i++ {
		result = result * nb
	}
	return result
}
