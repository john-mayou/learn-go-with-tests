package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"word": "meaning"}
		assertEqual(t, dict.Search("word"), "meaning")
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}
