package chain

import (
	"go-word-chain/src/dictionary"
	"testing"
)

func TestAreWordsAdjacent(t *testing.T) {
	var tests = map[string]struct {
		firstWord      string
		secondWord     string
		expectedResult bool
	}{
		"words-adjacent": {
			"bear",
			"beer",
			true,
		},
		"words-not-adjacent": {
			"far",
			"cat",
			false,
		},
	}
	for keyTest, test := range tests {
		t.Run("Running chain test `"+keyTest+"`...", func(t *testing.T) {
			// Check if words are adjacent
			result := areWordsAdjacent(test.firstWord, test.secondWord)
			if test.expectedResult != result {
				t.Errorf("Test #%s failed; Received %t instead of %t", keyTest, result, test.expectedResult)
			}
		})
	}
}

func TestGetAdjacentWords(t *testing.T) {
	t.Run("Running chain test `get-adjacent-words`...", func(t *testing.T) {
		// Initialize dictionary
		dict := &dictionary.Dictionary{
			Words: []string{"car", "bat", "fan", "step", "call", "rat"},
		}
		word := "cat"
		// Retrieve adjacent words
		result := getAdjacentWords(word, dict)
		if len(result) != 4 {
			t.Errorf("Test `get-adjacent-words` failed")

		}
	})
}

func TestRunBFS(t *testing.T) {
	t.Run("Running chain test `run-bfs`...", func(t *testing.T) {
		// Initialize dictionary
		dict := &dictionary.Dictionary{
			Words: []string{"cat", "bat", "dot", "lag", "fix", "dat", "dog"},
		}
		startWord := "cat"
		endWord := "dog"
		// Run breadth-first search to get fron startWord to endWord
		result := runBFS(startWord, endWord, dict)
		if len(result) != 5 {
			t.Errorf("Test `run-bfs` failed")

		}
	})
}

func TestBuildChain(t *testing.T) {
	t.Run("Running chain test `build-chain`...", func(t *testing.T) {
		// Initialize dictionary
		dict := &dictionary.Dictionary{
			Words: []string{"cat", "bat", "dot", "lag", "fix", "dat", "dog"},
		}
		startWord := "cat"
		endWord := "dog"
		result := BuildChain(startWord, endWord, dict)

		// Expected result
		expResult := []string{"cat", "dat", "dot", "dog"}
		for key, orderedEntry := range result {
			if expResult[key] != orderedEntry {
				t.Errorf("Test `build-chain` failed")

			}
		}
	})
}

func TestBuildChainSingle(t *testing.T) {
	t.Run("Running chain test `build-chain-single`...", func(t *testing.T) {
		// Initialize dictionary
		dict := &dictionary.Dictionary{
			Words: []string{"cat", "bat", "dot", "lag", "fix", "dat", "dog"},
		}
		// startWord is the same with endWord
		startWord := "cat"
		endWord := "cat"
		result := BuildChain(startWord, endWord, dict)

		// Expected result
		expResult := []string{"cat"}
		for key, orderedEntry := range result {
			if expResult[key] != orderedEntry {
				t.Errorf("Test `build-chain` failed")

			}
		}
	})
}
