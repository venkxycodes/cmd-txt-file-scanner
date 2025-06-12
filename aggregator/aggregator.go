package aggregator

import "cmd-txt-file-scanner/word-counter"

type Aggregator struct {
	WordCounter word_counter.WordCounter
}

func NewAggregator(w word_counter.WordCounter) *Aggregator {
	return &Aggregator{WordCounter: w}
}

func (a *Aggregator) AggregateWordCounts(paths []string) (map[string]int64, error) {
	totalCounts := make(map[string]int64)

	for _, path := range paths {
		wordCounts, err := a.WordCounter.CountWords(path)
		if err != nil {
			return nil, err
		}
		for word, count := range wordCounts {
			totalCounts[word] += int64(count)
		}
	}
	return totalCounts, nil
}
