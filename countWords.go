package main

// func countWords(filename string) {
// 	file := fileToString(filename)
// 	println(*file)
// }

import (
	"sort"
	"strings"
)

type tuple struct {
	word  string
	count int
}

func sortWords(words map[string]int) []tuple {
	sortedWords := make([]tuple, 0)
	for word, count := range words {
		sortedWords = append(sortedWords, tuple{word, count})
	}

	// Sort the slice based on values
	sort.Slice(sortedWords, func(i, j int) bool {
		return sortedWords[i].count > sortedWords[j].count
	})

	return sortedWords
}

func countWords(filename string) map[string]int {
	input := *fileToString(filename)

	// Split the input string into words using whitespace as the delimiter
	words := strings.Fields(input)

	// Create a map to store word counts
	wordCounts := make(map[string]int)

	// Count the occurrences of each word
	for _, word := range words {
		wordCounts[word]++
	}

	return wordCounts
}
