package kata

import "strings"

func High(s string) string {
	// your code here
	words := strings.Fields(s)

	bestWord := ""
	bestScore := 0

	for _, word := range words {
		score := 0

		for _, ch := range word {
			score += int(ch - 'a' + 1)
		}

		if score > bestScore {
			bestScore = score
			bestWord = word
		}
	}

	return bestWord
}
