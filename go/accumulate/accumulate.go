package accumulate

// Accumulate takes a list of strings and applies a function to each value.
func Accumulate(in []string, converter func(string) string) []string {
	out := make([]string, len(in))
	for i, val := range in {
		out[i] = converter(val)
	}
	return out
}