package prime

import (
	"math"
)

//Nth when given a number n, determines what the nth prime is.
func Nth(nPrime int) (int, bool) {
	if nPrime <= 0 {
		return -1, false
	}
	counter := 1
	currentPrime := 2
	for i := 3; i > 0; i++ {
		if nPrime == counter {
			return currentPrime, true
		}
		if isPrime(i) {
			currentPrime = i
			counter++
		}

	}
	return -1, true
}

func isPrime(n int) bool {
	squareRoot := int(math.Ceil(math.Sqrt(float64(n))))
	for i := 2; i <= squareRoot; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
