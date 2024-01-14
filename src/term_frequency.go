package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func getStopWords() map[string]bool {
	stopWords, err := os.ReadFile("../stop_words.txt")
	if err != nil {
		log.Fatal(err)
	}
	stopWordsList := strings.Split(string(stopWords), ",")
	stopWordsMap := make(map[string]bool)
	for idx := 0; idx < len(stopWordsList); idx++ {
		stopWordsMap[stopWordsList[idx]] = true
	}
	return stopWordsMap
}

func processFile(inputFile string) string {
	content, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	contentString := strings.ToLower(string(content))
	return contentString
}

func replaceNonAlphanumeric(input string) string {
	reg := regexp.MustCompile("[^a-z0-9]")
	return reg.ReplaceAllString(input, " ")
}

func filterStopWords(input string, stopWords map[string]bool) []string {
	var processed []string
	words := strings.Split(input, " ")
	for idx := 0; idx < len(words); idx++ {
		word := words[idx]
		_, exists := stopWords[word]
		if !exists && len(word) >= 2 {
			processed = append(processed, word)
		}
	}
	return processed
}

func processAndFilter(content string, stopWords map[string]bool) []string {
	content = strings.ToLower(content)
	processedInput := replaceNonAlphanumeric(content)
	processedAndFiltered := filterStopWords(processedInput, stopWords)
	return processedAndFiltered
}

func calcFrequencies(words []string) map[string]int {
	var frequencies = make(map[string]int)
	for idx := 0; idx < len(words); idx++ {
		word := words[idx]
		frequencies[word] += 1
	}
	return frequencies
}

func printTopNFrequencies(frequencies map[string]int, n int) {
	keys := make([]string, 0, len(frequencies))
	for key := range frequencies {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return frequencies[keys[i]] < frequencies[keys[j]]
	})

	// fmt.Println(keys)

	lastN := keys[len(keys)-n:]

	for idx := n - 1; idx >= 0; idx-- {
		fmt.Println(lastN[idx], "-", frequencies[lastN[idx]])
	}
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal()
	}
	inputFile := os.Args[1]
	stopWords := getStopWords()
	content := processFile(inputFile)
	words := processAndFilter(content, stopWords)
	frequencies := calcFrequencies(words)
	printTopNFrequencies(frequencies, 25)
	// fmt.Println(frequencies)
}
