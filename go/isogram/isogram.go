package isogram

import (
	"strings"
	"unicode"
)

func IsIsogram(input string) bool {
	lower := strings.ToLower(input)
	charMap := make(map[rune]int)
	for _, c := range lower {
		if unicode.IsLetter(c) {
			if _, ok := charMap[c]; ok {
				return false
			}
		}
		charMap[c] = 1
	}
	return true
}
