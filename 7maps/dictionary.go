package maps

import "errors"

// Maps allow you to store items in a manner similar to a dictionary. You can think of
// the key as the word and the value as the definition. And what better way is there to
// learn about Maps than to build our own dictionary?

type Dictionary map[string]string

var (
	ErrNotFound = errors.New("could not find the word you were looking for")
	ErrorWordExists = errors.New("cannot add word because it already exists")
	ErrorWordDeosNotExists = errors.New("cannot perform operation on word because it does not exist")
)

// In order to make this pass, we are using an interesting property of the map lookup, 
// It can return 2 values. The second value is a boolean which indicates if the key was
// found successfully. This property allows us to differentiate between a word that doesn't
// exist and a word that just doesn't have a definition.

func (d Dictionary) Search( word string) (string, error){
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error{
	_, err := d.Search(word)

	switch err{
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrorWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error{

	_, err := d.Search(word)

	switch err{
	case ErrNotFound:
		return ErrorWordDeosNotExists
	case nil:
		d[word] = definition
	default:
		return err
	}
	
	return nil
}

func (d Dictionary) Delete(word string) error{
	_, err := d.Search(word)

	switch err{
	case ErrNotFound:
		return ErrorWordDeosNotExists
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}

type DictionaryErr string

func (e DictionaryErr) Error() string{
	return string(e)
}