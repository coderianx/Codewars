package kata

func GetCount(str string) (count int) {
	// Enter solution here
	count = 0

	for _, char := range str {
		if char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			count++
		}
	}

	return count
}
