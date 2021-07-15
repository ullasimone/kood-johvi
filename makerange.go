/*Write a function that takes an int min and an int max as parameters. That function returns a slice of int with all the values between min and max.
Min is included, and max is excluded.
If min is superior or equal to max, a nil slice is returned.
append is not allowed for this exercise.*/

package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	size := max - min
	answer := make([]int, size)

	for i := 0; i < size; i++ {
		answer[i] = min + i
	}
	return answer
}
