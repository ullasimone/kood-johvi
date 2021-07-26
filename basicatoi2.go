package piscine

func BasicAtoi2(s string) int {
	res := 0     // final answer
	counter := 0 // counts chars
	checker := true
	a_s := []rune(s)         // makes string into runes
	for _, ch := range a_s { // searches all a_s chars
		if ch >= 48 && ch <= 57 { // checks if ch is larger than '0' and smaller than '9'
			for i := '0'; i < ch; i++ {
				counter++
			}
			res = res*10 + counter // matemaagia: adds chars from loop to each other
			counter = 0
		} else { // checks if char != int
			checker = false
		}
	}
	if checker { // if char == int returns res
		return res
	} else { // if char != int returns 0
		return 0
	}
}
