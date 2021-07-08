// Write a function that takes two pointers to an int (*int) and swaps their contents.

package piscine

func Swap(a *int, b *int) {
	asdfA := *a
	asdfB := *b
	*a = asdfA
	*b = asdfB
}
