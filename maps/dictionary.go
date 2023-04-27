package maps

import (
	"errors"
)

var (
	ErrNotFound   = errors.New("could not find value")
	ErrWordExists = errors.New("word already exists in the dictionary")
)

type Dictionary map[string]string

type Searcher interface {
	Search(string) (string, error)
}

type Adder interface {
	Add(string, string) error
}

func (d Dictionary) Search(word string) (string, error) {
	data, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return data, nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}
