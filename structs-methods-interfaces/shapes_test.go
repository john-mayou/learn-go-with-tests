package shapes

import (
	"fmt"
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	tests := []struct {
		shape    Shape
		expected float64
	}{
		{shape: &Rectangle{Width: 3.0, Height: 4.0}, expected: 14.0},
		{shape: &Rectangle{Width: 10.0, Height: 10.0}, expected: 40.0},
		{shape: &Circle{Radius: 5.0}, expected: 31.4159},
		{shape: &Circle{Radius: 10.0}, expected: 62.8319},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("returns expected perimeter for shape %#v", tt.shape)
		t.Run(testname, func(t *testing.T) {
			actual := RoundFloat(tt.shape.Perimeter(), 4)
			if actual != tt.expected {
				t.Errorf("expected %v but got %v", tt.expected, actual)
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
		{shape: &Circle{Radius: 5.0}, expected: 78.5398},
		{shape: &Circle{Radius: 10.0}, expected: 314.1593},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("returns expected area for shape %#v", tt.shape)
		t.Run(testname, func(t *testing.T) {
			actual := RoundFloat(tt.shape.Area(), 4)
			if actual != tt.expected {
				t.Errorf("expected %v but got %v", tt.expected, actual)
			}
		})
	}
}

func TestRoundFloat(t *testing.T) {
	tests := []struct {
		float    float64
		places   int
		expected float64
	}{
		{float: 0.11111, places: 0, expected: 0.0},
		{float: 0.11111, places: 1, expected: 0.1},
		{float: 0.11111, places: 2, expected: 0.11},
		{float: 0.11111, places: 3, expected: 0.111},
		{float: 0.11111, places: 4, expected: 0.1111},
		{float: 0.11115, places: 4, expected: 0.1112},
		{float: 0.11119, places: 4, expected: 0.1112},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("rounds float %v %v places correctly", tt.float, tt.places)
		t.Run(testname, func(t *testing.T) {
			actual := RoundFloat(tt.float, tt.places)
			if actual != tt.expected {
				t.Errorf("expected %v but got %v", tt.expected, actual)
			}
		})
	}
}

func RoundFloat(float float64, places int) float64 {
	multiplier := math.Pow10(places)
	return math.Round(float*multiplier) / multiplier
}
