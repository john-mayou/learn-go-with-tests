package dictionary

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	meaning, found := d[word]
	if !found {
		return "", ErrNotFound
	}
	return meaning, nil
}

func (d Dictionary) Add(word, meaning string) {
	d[word] = meaning
}
