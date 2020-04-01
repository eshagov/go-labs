package hw_3

import (
	"regexp"
	"sort"
	"strings"
)


func TopTenWords(text string) []string {
	words := splitText(text)
	wordCounts := make(map[string]int)
	for i := 0; i < len(words); i++ {
		wordCounts[words[i]] = wordCounts[words[i]] + 1
	}
	topTenWords := topWords(wordCounts)

	return topTenWords
}

func splitText(text string) []string {
	words := make([]string, 0)
	var currentWord strings.Builder
	var IsLetter = regexp.MustCompile(`^[a-zA-Z0-9]+$`).MatchString

	for _, symbol := range text {
		if IsLetter(string(symbol)) {
			currentWord.WriteRune(symbol)
		} else if currentWord.Len() > 0 {
			words = append(words, currentWord.String())
			currentWord.Reset()
		}
	}

	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
		currentWord.Reset()
	}

	return words
}

func topWords(unsortedMap map[string]int) []string {
	topWords := make([]string, 0)
	keys := make([]string, 0, len(unsortedMap))
	for key := range unsortedMap {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool { return unsortedMap[keys[i]] > unsortedMap[keys[j]] })

	for index, key := range keys {
		topWords = append(topWords, key)
		if index == 9 {
			break
		}
	}

	return topWords
}
