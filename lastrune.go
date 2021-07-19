package piscine

func LastRune(s string) rune {
	lR := []rune(s)      // changes string into runes
	return lR[len(lR)-1] // returns strings last rune as len(s)-1
}
