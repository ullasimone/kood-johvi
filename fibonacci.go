package piscine

// import "fmt"

func Fibonacci(index int) int {
	if index <= 0 {
		return index
	} else {
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}

/*func main() {
	arg1 := 4
	fmt.Println(Fibonacci(arg1))
}*/
