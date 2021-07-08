package piscine

func UltimateDivMod(a *int, b *int) {
	var numA := *a / *b
	var numB := *a % *b
	*a = numA
	*b = numB
}
