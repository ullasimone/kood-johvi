package piscine

func Swap(a *int, b *int) { // two inputs
	tempA := *a // var tempA pointer
	tempB := *b // var tempB pointer
	*a = tempB  // swaps places
	*b = tempA  // swaps places
}

// pointer refers to specified block in memory which has certain queue number
