package chain

import (
	"go-word-chain/src/dictionary"
	"go-word-chain/src/queue"
)

// BuildChain tries to find a path between startWord and endWord and returns it
func BuildChain(startWord string, endWord string, dict *dictionary.Dictionary) []string {
	// Edge case - equal strings
	if startWord == endWord {
		return []string{startWord}
	}

	// Filter dictionary by length
	dict = dict.FilterByLength(len(startWord))

	// Treat the words and relations between them as a tree
	// Run a BFS to get from the startWord (root) to the endWord (leaf)
	paths := runBFS(startWord, endWord, dict)

	// Get the path by traversing the tree in reverse order, from the endWord to the startWord
	word := endWord
	chain := make([]string, 0)
	for word != startWord {
		chain = append([]string{word}, chain...)
		word = paths[word]
	}
	chain = append([]string{startWord}, chain...)

	return chain
}

// Checks if two words are adjacent - differ by only one character
func areWordsAdjacent(firstWord string, secondWord string) bool {
	diff := 0
	// Iterate over the word chars and find how many differences there are
	for i := 0; i < len(firstWord); i++ {
		if firstWord[i] != secondWord[i] {
			diff++
		}
	}

	// If there's only one different character, the words are adjacent
	if diff == 1 {
		return true
	}

	return false
}

// Gets all adjacent words for a given word
func getAdjacentWords(word string, dict *dictionary.Dictionary) []string {
	// Init adjacent word list
	adjWord := make([]string, 0)
	for _, entry := range dict.Words {
		if areWordsAdjacent(word, entry) {
			adjWord = append(adjWord, entry)
		}
	}

	return adjWord
}

// Run a Breadth-First-Search to find the endWord and return a paths map as a result
func runBFS(startWord string, endWord string, dict *dictionary.Dictionary) map[string]string {
	// Initialize paths map
	paths := make(map[string]string)
	// The tree will be represented using a queue
	q := &queue.Queue{}
	q.Enqueue(startWord)

	// Each iteration adds the adjacent words of the currently processed word to the queue
	for !q.IsEmpty() {
		// Process the first word in the queue
		elem := q.Dequeue()
		word := elem.Value.(string)
		if word == endWord {
			// Found the endWord
			break
		}
		// Get adjacent words
		adjWords := getAdjacentWords(word, dict)
		for _, adjWord := range adjWords {
			// Check if the adjacent word exists in the paths map
			if _, ok := paths[adjWord]; !ok {
				// If not, add it both to the paths map and to the queue
				paths[adjWord] = word
				q.Enqueue(adjWord)
			}
		}
	}

	return paths
}
