package countdown

import (
	"bytes"
	"slices"
	"testing"
)

type SpyOperations struct {
	Calls []string
}

const (
	sleep = "sleep"
	write = "write"
)

func (s *SpyOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func TestCountdown(t *testing.T) {
	t.Run("counts down", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		sleeper := &SpyOperations{}
		Countdown(buffer, sleeper, 3, "Go!")
		assertEqual(t, buffer.String(), "3\n2\n1\nGo!")
	})
	t.Run("write before every sleep", func(t *testing.T) {
		spy := &SpyOperations{}
		Countdown(spy, spy, 3, "Go!")
		assertEqualSlices(t, spy.Calls, []string{write, sleep, write, sleep, write, sleep, write})
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func assertEqualSlices[V comparable](t testing.TB, got, want []V) {
	t.Helper()
	if !slices.Equal(got, want) {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
