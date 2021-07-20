/*Write a function IsSorted that returns true, if the slice of int is sorted, and that returns false otherwise.
The function passed in parameter returns a positive int if a (the first argument) is superior to b (the second argument), it returns 0 if they are equal and it returns a negative int otherwise.
To do your testing you have to write your own f function.*/

package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	if len(a) > 1 {
		if f(a[0], a[1]) >= 0 {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) < 0 {
					return false
				}
			}
		}
		if f(a[0], a[1]) <= 0 {
			for i := 0; i < len(a)-1; i++ {
				if f(a[i], a[i+1]) > 0 {
					return false
				}
			}
		}

	}
	return true
}
