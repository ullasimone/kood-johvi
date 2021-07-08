package main

func PointOne(n *int) {

	*n = *n + 1
}

func main() {

	n := 0
	PointOne(&n)

	Println(n)
}
