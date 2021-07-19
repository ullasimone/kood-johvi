package piscine

func Atoi(s string) int { // atoi changes the string numbers into int numbers
	o_number := 0    // final answer
	c := 0           // in loop this makes number into runes
	checker := true  // checks if string contains only numbers
	a_s := []rune(s) // innerstring that's converted into runes
	pl := 1          // ???

	if s != "" { // checks if string contains anything at all
		if byte(a_s[0]) == 45 { // checks if the first rune is '-'/negative
			pl = -1
			a_s[0] = '0' // converts the '-' into '0'/nothing
		} else if byte(a_s[0]) == 43 { // checks if the first rune is '+'/positive
			a_s[0] = '0'
		}
	}
	for _, word := range a_s { // looks for every rune (tähemärk) in string
		if word >= 48 && word <= 57 { // checks that every rune unicode is higher than 48 ('0') & smaller than 57 ('9')
			for i := '0'; i < word; i++ {
				c++ // if rune is a number, then the counter adds numbers to var i
			}
			o_number = o_number*10 + c // multiplying numbers to add numbers from string to each other
			c = 0
		} else {
			checker = false // if there isn't numbers the answer is '0'/nothing
		}
	}

	if checker { // if checker = true and there are numbers
		return o_number * pl // final answer * (-)1
	} else { // everything else returns 0
		return 0
	}
}
