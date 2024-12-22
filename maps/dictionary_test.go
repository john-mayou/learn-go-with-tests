package dictionary

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dict := Dictionary{"word": "meaning"}
		meaning, err := dict.Search("word")
		assertNoErr(t, err)
		assertEqual(t, meaning, "meaning")
	})
	t.Run("unknown word", func(t *testing.T) {
		dict := Dictionary{"word": "meaning"}
		_, err := dict.Search("random")
		assertErr(t, err, "unknown word")
	})
}

func assertEqual[V comparable](t testing.TB, got, want V) {
	t.Helper()
	if got != want {
		t.Errorf("got %v but wanted %v", got, want)
	}
}

func assertErr(t testing.TB, err error, msg string) {
	t.Helper()
	if err == nil {
		t.Fatal("wanted err but did not get one")
	}
	if err.Error() != msg {
		t.Errorf("got %q but wanted %q", err.Error(), msg)
	}
}

func assertNoErr(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("wanted no err but got %q", err.Error())
	}
}
