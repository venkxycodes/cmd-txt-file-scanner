package main

import (
	"cmd-txt-file-scanner/appcontext"
	"cmd-txt-file-scanner/domain"
	"cmd-txt-file-scanner/services/line_counter"
	"container/heap"
	"log"
	"os"
)

func execute(rootDir string) {
	paths, err := appcontext.GetFileScanner().Scan(rootDir)
	if err != nil {
		log.Fatalf("err-failed-to-scan-directory: %v", err)
		return
	}

	line_counter.CountAllLines(paths)

	directoryDetails := domain.TextFiles{
		Paths:         paths,
		NumberOfFiles: int64(len(paths)),
	}
	wordCountMap, aggErr := appcontext.GetWordCounter().AggregateWordCounts(directoryDetails.Paths)
	if aggErr != nil {
		log.Fatalf("err-failed-to-aggregate-word-counts: %v", err)
	}
	directoryDetails.WordCountPerFile = wordCountMap
	h := &domain.MinHeap{}
	heap.Init(h)
	for word, count := range wordCountMap {
		if h.Len() < 10 {
			heap.Push(h, domain.WordCount{
				Word:  word,
				Count: count,
			})
		} else if count > (*h)[0].Count {
			heap.Pop(h)
			heap.Push(h, domain.WordCount{
				Word:  word,
				Count: count,
			})
		}
	}

	var top10 []domain.WordCount
	for h.Len() > 0 {
		top10 = append(top10, heap.Pop(h).(domain.WordCount))
	}
	for i, j := 0, len(top10)-1; i < j; i, j = i+1, j-1 {
		top10[i], top10[j] = top10[j], top10[i]
	}
	log.Printf("info-top-10-words-by-count:")
	for _, wc := range top10 {
		log.Printf("%s: %d\n", wc.Word, wc.Count)
	}
	log.Printf("Processed %d files\n", len(paths))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <directory>", os.Args[0])
	}
	rootDir := os.Args[1]
	appcontext.Init()
	execute(rootDir)
}
