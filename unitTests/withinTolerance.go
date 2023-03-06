package unitTests

import "math"

// Fix issues when the result differs from the expected one by more than a trillionth of a fraction
func WithinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}

	d := math.Abs(a - b)

	if b == 0 {
		return d < e
	}

	return (d / math.Abs(b)) < e
}
