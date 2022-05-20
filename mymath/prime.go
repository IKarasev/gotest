package mymath

func IsPrime(n int) bool {
	if n < 0 {
		n = -n
	}

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
