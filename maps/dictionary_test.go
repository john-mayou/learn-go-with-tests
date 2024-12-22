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
	t.Run("does not override word", func(t *testing.T) {
		dict := Dictionary{"word": "old meaning"}
		err := dict.Add("word", "new meaning")
		assertErr(t, err, ErrWordExists.Error())
		meaning, err := dict.Search("word")
		assertNoErr(t, err)
		assertEqual(t, meaning, "old meaning")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updates word", func(t *testing.T) {
		dict := Dictionary{"word": "old meaning"}
		err := dict.Update("word", "new meaning")
		assertNoErr(t, err)
		meaning, err := dict.Search("word")
		assertNoErr(t, err)
		assertEqual(t, meaning, "new meaning")
	})
	t.Run("returns error on unknown word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Update("word", "new meaning")
		assertErr(t, err, ErrNotFound.Error())
		_, err = dict.Search("word")
		assertErr(t, err, ErrNotFound.Error())
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
