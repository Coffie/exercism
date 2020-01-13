package proverb

import "fmt"

const (
	DefaultPhrase = "For want of a %s the %s was lost."
	FinalPhrase = "And all for the want of a %s."
)

// Proverb builds a poem given a list of proverbs
func Proverb(rhyme []string) []string {
	length := len(rhyme)
	if length == 0 {
		return nil
	}
	poem := make([]string, length)
	poem[length-1] = fmt.Sprintf(FinalPhrase, rhyme[0])
	for i := 0; i < length - 1; i++ {
		poem[i] = fmt.Sprintf(DefaultPhrase, rhyme[i], rhyme[i+1])
	}
	return poem
}
