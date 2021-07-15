/*Write a function that separates the words of a string and puts them in a string slice.
The separators are spaces, tabs and newlines.*/

package piscine

func SplitWhiteSpaces(s string) []string {
	var arrayString []string
	countWords := 1
	len := 0
	answer := ""
	for a := range s {
		if isWhiteSpace(s[a]) {
			countWords++
		}
		len++
	}
	arrayString = make([]string, countWords)
	i := 0
	for j, a := range s {
		if j+1 == len {
			arrayString[i] = answer + string(s[j])
		}
		if isWhiteSpace(s[j]) {
			if i <= countWords {
				arrayString[i] = answer
				i++
				answer = ""
			}
		} else {
			answer = answer + string(a)
		}
	}
	return arrayString
}

func isWhiteSpace(r byte) bool {
	if r == ' ' || r == '\n' {
		return true
	}
	return false
}
