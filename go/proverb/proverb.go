// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

const (
	DefaultPhrase = "For want of a %s the %s was lost."
	FinalPhrase = "And all for the want of a %s."
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	var poem []string
	var endPhrase string
	length := len(rhyme)
	if length > 0 {
		endPhrase = fmt.Sprintf(FinalPhrase, rhyme[0])
	}
	switch length {
	case 0:
		return poem
	case 1:
		return append(poem, endPhrase)
	default:
		for i, r := range rhyme {
			if i == length - 2 {
				poem = append(poem, fmt.Sprintf(DefaultPhrase, r, rhyme[i+1]))
				break
			}
			poem = append(poem, fmt.Sprintf(DefaultPhrase, r, rhyme[i+1]))
		}
	}
	return append(poem, endPhrase)
}
