/* Write a function that returns true if the int passed as parameter is a prime number. Otherwise it returns false.
The function must be optimized in order to avoid time-outs with the tester.
(We consider that only positive numbers can be prime numbers)
(We also consider that 1 is not a prime number)*/

package piscine

// import "fmt"

func IsPrime(nb int) bool {
	// optimize below 2
	if nb < 2 {
		return false
	} // loop - start to search
	for i := 2; i*i <= nb; i++ {
		// divide with i, if it has decimal, not a prime number
		if nb%i == 0 {
			return false
		}
	}
	return true
}
