package maps

import "errors"

var ErrNotFound = errors.New("could not find value")

type Dictionary map[string]string

type Searcher interface {
	Search(string) (string, error)
}

func (d Dictionary) Search(word string) (string, error) {
	data, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return data, nil
}
