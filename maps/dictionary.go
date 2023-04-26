package maps

import "errors"

type Dictionary map[string]string

type Searcher interface {
	Search(string) (string, error)
}

func (d Dictionary) Search(word string) (string, error) {
	data, ok := d[word]

	if !ok {
		return "", errors.New("missing value")
	}

	return data, nil
}
