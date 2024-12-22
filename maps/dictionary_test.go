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
		assertErr(t, err, ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		dict := Dictionary{}
		dict.Add("word", "meaning")
		meaning, err := dict.Search("word")
		assertNoErr(t, err)
		assertEqual(t, meaning, "meaning")
	})
	t.Run("override word", func(t *testing.T) {
		dict := Dictionary{"word": "old meaning"}
		err := dict.Add("word", "new meaning")
		assertErr(t, err, ErrWordExists.Error())
		meaning, err := dict.Search("word")
		assertNoErr(t, err)
		assertEqual(t, meaning, "old meaning")
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
