package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"go-word-chain/src/chain"
	"go-word-chain/src/dictionary"
)

func main() {
	// Check if required args are set
	cliArgs := os.Args
	if len(cliArgs) < 3 {
		log.Fatal(errors.New("Not enough arguments"))
	}
	startWord := os.Args[1]
	endWord := os.Args[2]
	// Check if the words have the same size
	if len(startWord) != len(endWord) {
		log.Fatal(errors.New("Words are of different lengths"))
	}
	// Loading dictionary
	dictionary, err := dictionary.InitFromFile("./dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}
	// Check if both words exist in the dictionary
	if !dictionary.Contains(startWord) {
		log.Fatal(errors.New("Start word not in dictionary"))
	}
	if !dictionary.Contains(endWord) {
		log.Fatal(errors.New("End word not in dictionary"))
	}

	// Build the chain and print it
	c := chain.BuildChain(startWord, endWord, dictionary)
	for _, link := range c {
		fmt.Println(link)
	}
}
