package countdown

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	t.Run("counts down", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, 3, "Go!")
		assertEqual(t, buffer.String(), "3\n2\n1\nGo!")
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
