package kata

func SquareSum(numbers []int) int {
	// your code here
	sum := 0

	for _, number := range numbers {
		sum += number * number
	}

	return sum
}
