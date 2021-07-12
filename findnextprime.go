/*Write a function that returns the first prime number that is equal or superior to the int passed as parameter.
The function must be optimized in order to avoid time-outs with the tester.
(We consider that only positive numbers can be prime numbers)*/

package piscine

func FindNextPrime(nb int) int {
	primeNb := true

	for IsPrime(nb) {
		if IsPrime(nb) == primeNb {
			return nb
		}
		for IsPrime(nb) != primeNb {
			return nb
		}
	}
	return 0
}
