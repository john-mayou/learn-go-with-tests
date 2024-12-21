package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("returns sum of slice elements", func(t *testing.T) {
		assertEqualInt(t, Sum([]int{1, 2, 3, 4, 5}), 15)
		assertEqualInt(t, Sum([]int{1, 2, 3, 4, 6}), 16)
	})
}

func assertEqualInt(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
