package dictionary

import (
	"errors"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// Dictionary is used to store the list of words
type Dictionary struct {
	Words []string
}

// InitFromFile opens a dictionary file and returns it's content
func InitFromFile(filepath string) (d *Dictionary, err error) {
	// Check dictionary filepath not empty
	if len(filepath) == 0 {
		return nil, errors.New("cannot open config file")
	}

	// Open dictionary file
	dFile, err := os.Open(filepath)
	if err != nil {
		return nil, err

	}
	defer dFile.Close()
	return readInput(dFile)
}

func readInput(r io.Reader) (dict *Dictionary, err error) {
	// Read dictionary file
	rawDictionary, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	dict = &Dictionary{Words: strings.Split(string(rawDictionary), "\n")}

	return dict, nil
}

// Contains checks if a word is present in the dictionary
func (d *Dictionary) Contains(word string) bool {
	for _, entry := range d.Words {
		if entry == word {
			return true
		}
	}

	return false
}

// FilterByLength returns a dictionary with words of a specific length, based on the original dictionary
func (d *Dictionary) FilterByLength(length int) *Dictionary {
	if length == 0 {
		// Empty dictionary - invalid
		return nil
	}
	dict := &Dictionary{}
	// Find words of size `length`
	for _, entry := range d.Words {
		if len(entry) != length {
			continue
		}
		dict.Words = append(dict.Words, entry)
	}

	return dict
}
