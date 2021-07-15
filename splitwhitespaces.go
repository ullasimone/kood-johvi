/*Write a function that separates the words of a string and puts them in a string slice.
The separators are spaces, tabs and newlines.*/

package piscine

func SplitWhiteSpaces(s string) []string {
	a := 0
	noWhiteSpaces := false
	for i := range s {
		if noWhiteSpaces && i != 0 && (s[i-1] == ' ' || s[i-1] == '\n' || s[i-1] == '\t') && s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
			a++
		}
		if s[i] != ' ' && s[i] != '\n' && s[i] != '\t' {
			noWhiteSpaces = true
		}
	}
	a++

	b := 0
	answer := make([]string, a)
	oa := true
	for _, c := range s {
		if c == ' ' || c == '\n' || c == '\t' {
			if !oa {
				b++
			}
			oa = true
			continue
		}
		answer[b] = answer[b] + string(c)
		oa = false
	}
	return answer
}
