package triangle

import (
	"math"
)

type Kind int

const (
    NaT = iota
    Equ
    Iso
    Sca
)

// KindFromSides gets three values and evaluates which kind of triangle it makes
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	ab, ac, bc := a + b, a + c, b + c
	if invalidNumber(a) || invalidNumber(b) || invalidNumber(c) {
		k = NaT
	} else if ab < c || ac < b || bc < a {
		k = NaT
	} else if a == b && b == c {
		k = Equ
	} else if a == b || a == c || b == c {
		k = Iso
	} else {
		k = Sca
	}
	return k
}

// invalidNumber checks if a value is a valid side of a triangle
func invalidNumber(num float64) bool {
	if math.IsNaN(num) || math.IsInf(num, 1) || math.IsInf(num, -1) || num == 0.0 {
		return true
	}
	return false
}
