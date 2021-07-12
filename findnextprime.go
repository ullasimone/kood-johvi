/*Write a function that returns the first prime number that is equal or superior to the int passed as parameter.
The function must be optimized in order to avoid time-outs with the tester.
(We consider that only positive numbers can be prime numbers)*/

package piscine

/*import (
	"fmt"
)*/

func kapsas(nb int) bool {
	// optimize below 2
	if nb < 2 {
		return false
	} // loop - start to search
	for i := 2; i < nb; i++ {
		// divide with i, if it has decimal, not a prime number
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}
	for i := nb; ; i++ {
		if kapsas(i) {
			return i
		}
	}
}

/*func main() {
	fmt.Println(FindNextPrime(16))
	fmt.Println(FindNextPrime(7))
}*/
