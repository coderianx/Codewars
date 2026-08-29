package kata

func BouncingBall(h, bounce, window float64) int {
	// your code
	if h <= 0 || bounce <= 0 || bounce >= 1 || window >= h {
		return -1
	}

	count := 1

	h = h * bounce

	for h > window {
		count = count + 2

		h = h * bounce
	}

	return count
}
