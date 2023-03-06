package unitTests

import (
	"testUnitTest/cmd"
	"testing"
)

func TestDivide(t *testing.T) {
	// Arrange
	var x float64 = 18.55
	var y float64 = 5.3
	var expected float64 = 3.5

	// Act
	result := cmd.Divide(x, y)

	// Assert
	if !WithinTolerance(expected, result, 1e-12) {
		t.Errorf("Incorrect result. Expected %f, got %f", expected, result)
	}
}
