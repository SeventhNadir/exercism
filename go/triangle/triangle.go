package triangle

import "math"

//Kind of triangles you will find
type Kind int

//Categories of triangles
const (
	NaT = iota // not a triangle
	Equ        // equilateral
	Iso        // isosceles
	Sca        // scalene
)

// KindFromSides should have a comment documenting it.
func KindFromSides(a, b, c float64) Kind {
	if invalidTriangle(a, b, c) {
		return NaT
	}
	if a == b && b == c {
		return Equ
	}
	if a == b || b == c || a == c {
		return Iso
	}
	return Sca

}

func invalidTriangle(a, b, c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return true
	}

	if math.IsInf(a+b+c, 0) || math.IsNaN(a+b+c) {
		return true
	}

	if a+b < c || a+c < b || c+b < a {
		return true
	}
	return false
}
