package isogram

import (
	"strings"
)

func IsIsogram(input string) bool {
	lower := strings.ToLower(input)
	charMap := make(map[string]int)
	for _, c := range lower {
		if _, ok := charMap[string(c)]; ok {
			if string(c) == "-" || string(c) == " " {
				break
			}
			return false
		}
		charMap[string(c)] = 1
	}
	return true
}
