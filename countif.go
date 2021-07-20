// Write a function CountIf that returns, the number of elements of a string slice, for which the f function returns true.

package piscine

func CountIf(f func(string) bool, tab []string) int {
	ans := 0
	for _, res := range tab {
		if f(res) {
			ans = ans + 1
		}
	}
	return ans
}
