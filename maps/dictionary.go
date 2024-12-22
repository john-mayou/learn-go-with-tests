package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	meaning, found := d[word]
	if !found {
		return "", errors.New("unknown word")
	}
	return meaning, nil
}
