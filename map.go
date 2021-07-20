// Write a function Map that, for an int slice, applies a function of this type func(int) bool on each elements of that slice and returns a slice of all the return values.

package piscine

func Map(f func(int) bool, a []int) []bool {
	ans := make([]bool, len(a))
	for i, res := range a {
		ans[i] = f(res)
	}
	return ans
}
