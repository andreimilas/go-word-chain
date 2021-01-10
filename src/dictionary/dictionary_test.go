package dictionary

import (
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	t.Run("Running dictionary test `read-file`...", func(t *testing.T) {
		// Load the dictionary
		_, err := readInput(strings.NewReader("aberdeen\naccounting\naccuracy\naccurate\naccurately\naccused"))
		if err != nil {
			t.Errorf("Test `read-file` failed; Error: %s", err.Error())
		}
	})
}

func TestContains(t *testing.T) {
	var tests = map[string]struct {
		dict           *Dictionary
		word           string
		expectedResult bool
	}{
		"contains": {
			&Dictionary{Words: []string{"aberdeen", "accounting", "accuracy", "accurate", "accurately", "accused"}},
			"accounting",
			true,
		},
		"not-contains": {
			&Dictionary{Words: []string{"aberdeen", "accounting", "accuracy", "accurate", "accurately", "accused"}},
			"ascension",
			false,
		},
	}
	for keyTest, test := range tests {
		t.Run("Running dictionary test `"+keyTest+"`...", func(t *testing.T) {
			// Check if the dictionary contains the test word
			result := test.dict.Contains(test.word)
			if test.expectedResult != result {
				t.Errorf("Test #%s failed; Received %t instead of %t", keyTest, result, test.expectedResult)
			}
		})
	}
}

func TestFilterByLength(t *testing.T) {
	t.Run("Running dictionary test `filter-by-length`...", func(t *testing.T) {
		// Return a filtered dictionary by a specified length
		dict := &Dictionary{Words: []string{"a", "aa", "aaa", "aaron", "ab"}}
		result := dict.FilterByLength(2)
		if len(result.Words) != 2 {
			t.Errorf("Test `filter-by-length` failed; Received %d instead of %d", len(result.Words), 2)
		}
	})
}

func TestFilterByLength0(t *testing.T) {
	t.Run("Running dictionary test `filter-by-length-0`...", func(t *testing.T) {
		// Return a nil dictionary if the length filter is 0
		dict := &Dictionary{Words: []string{"a", "aa", "aaa", "aaron", "ab"}}
		result := dict.FilterByLength(0)
		if result != nil {
			t.Errorf("Test `filter-by-length-0` failed; Dictionary should be nil")
		}
	})
}
