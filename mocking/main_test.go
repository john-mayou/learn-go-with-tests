package main

import (
	"testing"
	"time"
)

type SpyTime struct {
	slept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.slept = duration
}

func TestConfigurableSleeper(t *testing.T) {
	t.Run("sleeps for configured duration", func(t *testing.T) {
		sleepTime := 5 * time.Second
		spyTime := &SpyTime{}
		sleeper := ConfigurableSleeper{duration: sleepTime, sleep: spyTime.Sleep}
		sleeper.Sleep()
		assertEqual(t, spyTime.slept, 5*time.Second)
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
