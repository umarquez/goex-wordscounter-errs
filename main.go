package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"words_counter/services/random_text_api"
	"words_counter/services/static_text"
)

func extractWords(txtGenerator func() (string, error)) []string {
	text, err := txtGenerator()
	if err != nil {
		err = fmt.Errorf(
			"text generator fails, impossible to continue: %v", err,
		)
		log.Fatal(err)
	}

	words := regexp.MustCompile(`[a-zA-Z/-]+`).FindAllString(text, -1)
	fmt.Printf("total words: %v\n\n", len(words))
	return words
}

func countWordsOccurrences(words ...string) (wordsCounters map[string]int) {
	wordsCounters = make(map[string]int)
	for _, word := range words {
		wordsCounters[strings.ToLower(word)]++
	}

	return
}

func printMapContent(content map[string]int) {
	for key, val := range content {
		fmt.Printf("%v - %v\n", key, val)
	}
}

func main() {
	wordList := extractWords(static_text.GetText)
	wordsCounters := countWordsOccurrences(wordList...)
	printMapContent(wordsCounters)

	fmt.Println("\n------------------\n")
	// changing source
	// single line call
	printMapContent(countWordsOccurrences(extractWords(random_text_api.GetText)...))
}
