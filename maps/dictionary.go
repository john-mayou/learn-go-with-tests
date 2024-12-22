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

var ErrWordExists = errors.New("cannot add word because it already exists")

func (d Dictionary) Add(word, meaning string) error {
	if _, found := d[word]; found {
		return ErrWordExists
	}
	d[word] = meaning
	return nil
}
