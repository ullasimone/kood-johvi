/*Write a function that takes an int min and an int max as parameters. That function returns a slice of int with all the values between min and max.
Min is included, and max is excluded.
If min is superior or equal to max, a nil slice is returned.
make is not allowed for this exercise.*/

package piscine

func AppendRange(min, max int) []int {
	var answer []int

	for i := min; i < max; i++ {
		answer = append(answer, i)
	}
	if min >= max {
		return nil
	}
	return answer
}
