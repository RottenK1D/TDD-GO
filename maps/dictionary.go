package maps

const (
	ErrNotFound   = DictionaryErr("could not find value")
	ErrWordExists = DictionaryErr("word already exists in the dictionary")
)

type DictionaryErr string

type Dictionary map[string]string

type Searcher interface {
	Search(string) (string, error)
}

type Adder interface {
	Add(string, string) error
}

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	data, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return data, nil
}

func (d Dictionary) Add(word, definition string) error {
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

func (d Dictionary) Update(word, definition string) error {
	d[word] = definition
	return nil
}
