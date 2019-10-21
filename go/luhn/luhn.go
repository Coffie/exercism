package luhn

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

func Valid(value string) bool {
	// Trim all whitespaces
	// Alternate way, not as fast
	//trimmedValue := strings.Join(strings.Fields(value), "")
	trimmedValue := strings.ReplaceAll(value, " ", "")
	length := utf8.RuneCountInString(trimmedValue)
	if length < 2 {
		return false
	}
	sum := 0
	for pos, char := range trimmedValue {
		if !unicode.IsDigit(char) {
			return false
		}
		valInt := int(char) - 48
		switch length % 2 {
		case 0:
			if pos % 2 == 0 {
				val := valInt * 2
				if val >= 10 {
					val -= 9
				}
				sum += val
			} else {
				sum += valInt
			}
		case 1:
			if pos % 2 == 1 {
				val := valInt * 2
				if val >= 10 {
					val -= 9
				}
				sum += val
			} else {
				sum += valInt
			}
		}
	}
	if sum % 10 == 0 {
		return true
	}
	return false
}
