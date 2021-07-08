// Write a function that reorders a slice of int in ascending order.

package piscine

func SortIntegerTable(table []int) {
	a := len(table)
	for s := 0; s < a; s++ {
		for t := 0; t < a; t++ {
			if table[s] < table[t] {
				b := table[s]
				table[s] = table[t]
				table[j] = a
			}
		}
	}
}
