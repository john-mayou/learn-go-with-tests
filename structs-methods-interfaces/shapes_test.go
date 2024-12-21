package shapes

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	tests := []struct {
		width    float64
		height   float64
		expected float64
	}{
		{width: 3.0, height: 4.0, expected: 14.0},
		{width: 10.0, height: 10.0, expected: 40.0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("returns expected perimeter for width %v and height %v", tt.width, tt.height)
		t.Run(testname, func(t *testing.T) {
			actual := Perimeter(Rectangle{Width: tt.width, Height: tt.height})
			if actual != tt.expected {
				t.Errorf("expected %.2f but got %.2f", tt.expected, actual)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		width    float64
		height   float64
		expected float64
	}{
		{width: 3.0, height: 4.0, expected: 12.0},
		{width: 10.0, height: 10.0, expected: 100.0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("returns expected area for width %v and height %v", tt.width, tt.height)
		t.Run(testname, func(t *testing.T) {
			actual := Area(Rectangle{Width: tt.width, Height: tt.height})
			if actual != tt.expected {
				t.Errorf("expected %.2f but got %.2f", tt.expected, actual)
			}
		})
	}
}
