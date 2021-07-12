/*Write a function that returns the first prime number that is equal or superior to the int passed as parameter.
The function must be optimized in order to avoid time-outs with the tester.
(We consider that only positive numbers can be prime numbers)*/

package piscine

func IsPrime(nb int) bool {
	if nb < 2 {
		return false
	}
	for i := 2; i < nb; i++ {
		if nb%i == 0 {
			return false
		}
	}
	return true
}

func FindNextPrime(nb int) int {
	nextprime := nb - 1
	for i := nb + 1; i > nb; nextprime++ {
		if IsPrime(nextprime) {
			return nextprime
		}
		i++
	}
	return nextprime
}

/*func main() {
	fmt.Println(FindNextPrime(6))
	fmt.Println(FindNextPrime(9))
}*/
