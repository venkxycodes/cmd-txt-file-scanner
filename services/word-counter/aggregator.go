package word_counter

import (
	"log"
	"runtime"
	"sync"
)

func (wc *WordCounter) AggregateWordCounts(paths []string) (map[string]int64, error) {
	totalCounts := make(map[string]int64)
	var (
		mu      sync.Mutex
		wg      sync.WaitGroup
		jobs    = make(chan string)
		results = make(chan map[string]int64)
	)

	workerCount := runtime.NumCPU()
	for i := 0; i < workerCount; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for path := range jobs {
				counts, err := wc.CountWords(path)
				if err != nil {
					log.Printf("err-counting-words, path: %s", path)
					continue
				}
				results <- counts
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	go func() {
		for _, p := range paths {
			jobs <- p
		}
		close(jobs)
	}()

	for counts := range results {
		mu.Lock()
		for word, cnt := range counts {
			totalCounts[word] += cnt
		}
		mu.Unlock()
	}
	return totalCounts, nil
}
