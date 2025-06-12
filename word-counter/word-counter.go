package word_counter

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

type WordCounter struct{}

func NewWordCounter() *WordCounter {
	return &WordCounter{}
}

func (w *WordCounter) CountWords(path string) (map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	counts := make(map[string]int)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		for _, word := range tokenize(line) {
			counts[word]++
		}
	}
	return counts, nil
}

func tokenize(line string) []string {
	return strings.FieldsFunc(strings.ToLower(line), func(r rune) bool {
		return !unicode.IsLetter(r)
	})
}
