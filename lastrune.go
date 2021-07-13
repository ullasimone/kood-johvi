package piscine

func LastRune(s string) rune {
	lS := []rune(s)
	return lS[len(lS)-1]
}
