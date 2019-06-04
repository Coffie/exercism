// Package twofer implements a sharing of an arbitrary item
// between two people.
package twofer

import "fmt"

// ShareWith returns a string where two of something is
// splitted between 'name' and me
func ShareWith(name string) string {
	if name == "" {
		name = "you"
	}
	return fmt.Sprintf("One for %s, one for me.", name)
}
