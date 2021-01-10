# CodeKata - Kata19: Word Chains

## Author
Andrei Milas

## Description
Proposed solutions for Kata19 in Go (http://codekata.com/kata/kata19-word-chains/).

## Implementation
Given that each node in the dictionary has several adjacent words and each adjacent word has it's own collection of adjacent words, a tree structure was considered.
We start by adding the start word to the tree, as root, followed by each adjacent word until the end word is found. At this moment, a path exists between the two words, representing the word chain.

## Running the project
You can run the project with the following command
```
go run main.go <startWord> <endWord>
```
Where `startWord` and `endWord` are CLI arguments, marking the first and last elements of the chain.
