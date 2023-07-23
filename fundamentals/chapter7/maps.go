package maps

const (
	errNotFound         = DictionaryErr("word not found")
	errWordExists       = DictionaryErr("word already added")
	errWordDoesntExists = DictionaryErr("word doesn't exist")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

func Search(dictionary map[string]string, key string) string {
	return dictionary[key]
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	definition, ok := d[key]
	if !ok {
		return "", errNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		d[key] = value
	case nil:
		return errWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case errNotFound:
		return errWordDoesntExists
	case nil:
		d[key] = value
	default:
		return err
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
