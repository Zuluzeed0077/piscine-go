package piscine

func FindNextPrime(nb int) int {
	if nb <= 2 {
		return 2
	}

	// If nb is even and > 2, make it odd (primes > 2 are odd)
	if nb%2 == 0 {
		nb++
	}

	for {
		if isPrime(nb) {
			return nb
		}
		nb += 2
	}
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	if n == 2 || n == 3 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i*i <= n; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}
