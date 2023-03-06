package unitTests

import (
	"testUnitTest/cmd"
	"testing"
)

func TestDivide(t *testing.T) {
	// Arrange
	testTable := []struct {
		x        float64
		y        float64
		expected float64
	}{
		{
			x:        18.55,
			y:        5.3,
			expected: 3.5,
		},
		{
			x:        10,
			y:        5,
			expected: 2,
		},
		{
			x:        158645.784356,
			y:        578.6548954,
			expected: 274.16303848312695014314053273971,
		},
	}

	// Act
	for _, testCase := range testTable {
		result := cmd.Divide(testCase.x, testCase.y)

		// Run with "go test -v" to see logs
		t.Logf("Calling Divide(%f, %f), result: %f\n", testCase.x, testCase.y, testCase.expected)

		// Assert
		if !WithinTolerance(testCase.expected, result, 1e-12) {
			t.Errorf("Incorrect result. Expected %f, got %f", testCase.expected, result)
		}
	}
}
