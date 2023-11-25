package main

import (
	"sort"
	"strings"
)

type tuple struct {
	word  string
	count int
}

func sortAndFilterWords(words map[string]int) []tuple {
	sortedWords := make([]tuple, 0)
	for word, count := range words {
		if hasLetters(word) {
			sortedWords = append(sortedWords, tuple{word, count})
		}
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

func getWordsLonger(words []tuple, minLength int) []tuple {
	var result []tuple
	for _, v := range words {

		if isAcronym(v.word) || len([]rune(v.word)) >= minLength {

			result = append(result, v)
		}
	}
	return result
}

func getWordsWithMoreCount(words []tuple, minCount int) []tuple {
	var result []tuple
	for _, v := range words {
		if isAcronym(v.word) || v.count >= minCount {
			result = append(result, v)
		}
	}
	return result
}

func getWordsAfterFiltering(filename string, minLen int, minCount int) []tuple {
	all_words := countWords(filename)
	result := sortAndFilterWords(all_words)

	if minLen != 0 {
		result = getWordsLonger(result, minLen)
	}
	if minCount != 0 {
		result = getWordsWithMoreCount(result, minCount)
	}

	return result
}
