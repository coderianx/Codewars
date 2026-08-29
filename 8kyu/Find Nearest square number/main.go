package kata

import "math"

func NearestSq(n int) int {
	// Code goes here
	root := int(math.Sqrt(float64(n)))

	if root*root == n {
		return n
	}

	lower := root * root
	upper := (root + 1) * (root + 1)

	if n-lower < upper-n {
		return lower
	}

	return upper
}
