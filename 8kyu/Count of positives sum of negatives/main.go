package kata

func CountPositivesSumNegatives(numbers []int) []int {
	// your code here
	var res []int

	var positiveSum []int
	negativeSum := 0

	for _, num := range numbers {
		if num > 0 {
			positiveSum = append(positiveSum, num)
		} else {
			negativeSum += num
		}
	}

	res = append(res, len(positiveSum), negativeSum)

	return res
}
