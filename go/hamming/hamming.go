// Package hamming compares the difference between two
// DNA strands of equal length
package hamming

import "errors"

// Distance calculates the difference between two DNA
// strands, if they are of equal length. Returns error
// if length is not equal.
func Distance(a, b string) (int, error) {
	diff := 0
	if len(a) != len(b) {
		return diff, errors.New("DNA not of equal length")
	}
	for i, v := range a {
		if string(v) != string(b[i]) {
			diff++
		}
	}
	return diff, nil

}
