package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("returns string repeated 5 times", func(t *testing.T) {
		assertEqualStr(t, Repeat("a"), "aaaaa")
		assertEqualStr(t, Repeat("ab"), "ababababab")
	})
}

func assertEqualStr(t testing.TB, actual, expected string) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected %q but got %q", expected, actual)
	}
}
