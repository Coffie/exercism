package triangle


// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
    // Pick values for the following identifiers used by the test program.
    NaT = 0
    Equ = 1
    Iso = 2
    Sca = 3
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	input := [3]float64{a,b,c}
	sides := make(map[float64]int)
	for side := range input {
		if _, ok := sides[input[side]]; ok {
			sides[input[side]] += 1
		} else {
			sides[input[side]] = 1
		}
	}
	if _, ok := sides[0.0]; ok {
		k = NaT
	}
	
	return k
}
