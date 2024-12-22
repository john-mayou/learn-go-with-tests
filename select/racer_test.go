package racer

import "testing"

func TestRacer(t *testing.T) {
	t.Run("returns fastest url", func(t *testing.T) {
		slowURL := "http://www.facebook.com"
		fastURL := "http://www.quii.dev"
		assertEqual(t, Racer(slowURL, fastURL), fastURL)
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
