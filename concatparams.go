/*Write a function that takes the arguments received in parameters and returns them as a string. The string is the result of all the arguments concatenated with a newline (\n) between.*/

package piscine

func ConcatParams(args []string) string {
	var answer string
	for i, str := range args {
		if i > 0 {
			str = "\n" + str
		}
		answer = answer + str
	}
	return answer
}
