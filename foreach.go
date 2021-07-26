// Write a function ForEach that, for an int slice, applies a function on each elements of that slice.

package piscine

func ForEach(f func(int), a []int) {
	for _, res := range a {
		f(res)
	}
}
