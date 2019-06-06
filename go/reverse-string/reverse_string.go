// Package reverse implements a tool for reversing utf-8 strings
package reverse

// Reverse takes a string as input and returns the string in reverse
func Reverse(input string) (r string) {
	runeString := []rune(input)
	for i, j := 0, len(runeString) -1; i < j; i, j = i + 1, j - 1 {
		runeString[i], runeString[j] = runeString[j], runeString[i]
	}
	return string(runeString)
}
