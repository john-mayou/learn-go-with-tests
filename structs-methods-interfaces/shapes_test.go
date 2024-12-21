package shapes

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	tests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: &Rectangle{Width: 3.0, Height: 4.0}, expected: 14.0},
		{shape: &Rectangle{Width: 10.0, Height: 10.0}, expected: 40.0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("returns expected perimeter for shape %#v", tt.shape)
		t.Run(testname, func(t *testing.T) {
			actual := tt.shape.Perimeter()
			if actual != tt.expected {
				t.Errorf("expected %.2f but got %.2f", tt.expected, actual)
			}
		})
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: &Rectangle{Width: 3.0, Height: 4.0}, expected: 12.0},
		{shape: &Rectangle{Width: 10.0, Height: 10.0}, expected: 100.0},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("returns expected area for shape %#v", tt.shape)
		t.Run(testname, func(t *testing.T) {
			actual := tt.shape.Area()
			if actual != tt.expected {
				t.Errorf("expected %.2f but got %.2f", tt.expected, actual)
			}
		})
	}
}
