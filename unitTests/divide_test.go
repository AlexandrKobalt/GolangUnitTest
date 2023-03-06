package unitTests

import (
	"math"
	"testUnitTest/cmd"
	"testing"
)

func TestMain(t *testing.T) {
	// Arrange
	var x float64 = 18.55
	var y float64 = 5.3
	var expected float64 = 3.5

	// Act
	result := cmd.Divide(x, y)

	// Assert
	if !withinTolerance(expected, result, 1e-12) {
		t.Errorf("Incorrect result. Expected %f, got %f", expected, result)
	}
}

// Fix issues when the result differs from the expected one by more than a trillionth of a fraction
func withinTolerance(a, b, e float64) bool {
	if a == b {
		return true
	}

	d := math.Abs(a - b)

	if b == 0 {
		return d < e
	}

	return (d / math.Abs(b)) < e
}
