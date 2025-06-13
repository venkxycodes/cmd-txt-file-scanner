package line_counter

import (
	"cmd-txt-file-scanner/domain"
	"log"
	"runtime"
	"sync"
)

func CountAllLines(paths []string) domain.LinesInEachPath {
	var (
		jobs    = make(chan string, len(paths))
		results = make(chan map[string]int64)
	)
	workerCount := runtime.NumCPU()
	var wg sync.WaitGroup
	wg.Add(workerCount)

	for i := 0; i < workerCount; i++ {
		go func() {
			defer wg.Done()
			for path := range jobs {
				n, err := CountLines(path)
				if err != nil {
					log.Printf("error counting lines in %s: %v", path, err)
					continue
				}
				temp := map[string]int64{path: n}
				results <- temp
			}
		}()
	}

	for _, p := range paths {
		jobs <- p
	}
	close(jobs)

	go func() {
		wg.Wait()
		close(results)
	}()

	lineCountMap := make(domain.LinesInEachPath, len(paths))
	for counts := range results {
		for path, count := range counts {
			lineCountMap[path] += count
		}
	}
	return lineCountMap
}
