package integers

import "testing"

func TestAdder(t *testing.T) {
	t.Run("returns expected result", func(t *testing.T) {
		assertEqualInt(t, Add(2, 2), 4)
		assertEqualInt(t, Add(2, 3), 5)
	})
}

func assertEqualInt(t testing.TB, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %d but got %d", expected, actual)
	}
}
