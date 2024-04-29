package main

import "errors"

type Dict map[string]string

var ErrorNotFound = errors.New("could not find the word you were looking for")

func (d Dict) Search(word string) (string, error) {
	value, ok := d[word]

	if !ok {
		return "", ErrorNotFound
	}
	return value, nil
}

func (d Dict) Add(word, def string) {
	d[word] = def
}
