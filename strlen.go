package piscine

func StrLen(s string) int {
	counter := 0  // counter
	for range s { // search every character
		counter++ // count every character
	}
	return counter // string characters length (including spaces, numbers, symbols)
}
