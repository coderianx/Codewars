package kata

import "strings"

func AbbrevName(name string) string {
	words := strings.Fields(name)

	return strings.ToUpper(string(words[0][0])) + "." +
		strings.ToUpper(string(words[1][0]))
}
