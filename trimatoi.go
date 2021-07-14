// Write a function that transforms a number within a string, in an int.
// For this exercise the handling of the - sign has to be taken into account. If the sign is encountered before any number it should determine the sign of the returned int.
// This function will only return an int. In case of invalid input, the function should return 0.
// Note: There will never be more than one sign by string in the tests.

package piscine

func TrimAtoi(s string) int {
	var array []int
	result := 0
	minusIndex := -1
	firstDigitIndex := 0
	index := 0
	arrayCount := 0
	for _, rune := range s {
		if rune == '-' {
			minusIndex = index
		}
		if isDigit(rune) {
			if firstDigitIndex == 0 {
				firstDigitIndex = index
			}
			array = append(array, int(rune-'0'))
		}
		index++
	}

	for count := range array {
		arrayCount = count + 1
	}

	for i := 0; i < arrayCount; i++ {
		result = result*10 + array[i]
	}

	if minusIndex < firstDigitIndex && minusIndex != -1 {
		result = result * -1
	}
	return result
}

func isDigit(digit rune) bool {
	for a := '0'; a <= '9'; a++ {
		if digit == a {
			return true
		}
	}
	return false
}
