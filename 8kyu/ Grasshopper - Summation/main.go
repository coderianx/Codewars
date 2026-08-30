package kata

func Summation(n int) int {
	// the sleeper must awaken!
	sum := 0

	for i := 1; i <= n; i++ {
		sum = sum + i
	}

	return sum
}
