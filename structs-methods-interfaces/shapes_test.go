package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	assertEqualFloat64(t, Perimeter(3.0, 4.0), 14.0)
	assertEqualFloat64(t, Perimeter(10.0, 10.0), 40.0)
}

func TestArea(t *testing.T) {
	assertEqualFloat64(t, Area(3.0, 4.0), 12.0)
	assertEqualFloat64(t, Area(10.0, 10.0), 100.0)
}

func assertEqualFloat64(t testing.TB, actual, expected float64) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %.2f but got %.2f", expected, actual)
	}
}
