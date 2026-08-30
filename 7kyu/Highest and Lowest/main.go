package kata

import (
	"fmt"
	"strconv"
	"strings"
)

func HighAndLow(in string) string {
	// Code here or
	fields := strings.Fields(in)

	var numbers []int

	for _, field := range fields {
		n, _ := strconv.Atoi(field)
		numbers = append(numbers, n)
	}

	min := numbers[0]
	max := numbers[0]

	for _, num := range numbers {
		if num < min {
			min = num
		}

		if num > max {
			max = num
		}
	}

	return fmt.Sprintf("%s %s", strconv.Itoa(max), strconv.Itoa(min))

}
