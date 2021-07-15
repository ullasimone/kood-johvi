/*Write a function that takes an int min and an int max as parameters. That function returns a slice of int with all the values between min and max.
Min is included, and max is excluded.
If min is superior or equal to max, a nil slice is returned.
append is not allowed for this exercise.*/

package piscine

func MakeRange(min, max int) []int {
	answer := make([]int, max-min)
	for i := 0; i < (max - min); i++ {
		answer[i] = min + i
	}
	if min >= max {
		return nil
	}
	return answer
}
