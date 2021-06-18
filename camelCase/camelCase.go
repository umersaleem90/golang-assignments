package camelCase

import (
	"strings"
)

func NumberOfWords(input string) int {
	var count int = 1
	for i, c := range input {
		var isCapital bool = strings.ToUpper(string(c)) == string(c)
		if i == 0 && isCapital {
			count = 0
		}
		if isCapital {
			count += 1
		}
	}
	return count
}
