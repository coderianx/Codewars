package kata

import "strconv"

func StringToNumber(str string) int {
	// your code here
	number, _ := strconv.Atoi(str)

	return number
}
