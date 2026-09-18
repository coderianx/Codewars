package kata

func Invert(arr []int) []int {
	numbers := []int{}

	for _, number := range arr {
		numbers = append(numbers, number*-1)
	}

	return numbers
}
