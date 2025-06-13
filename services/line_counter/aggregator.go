package line_counter

import (
	"cmd-txt-file-scanner/domain"
	"log"
	"runtime"
	"sync"
)

func CountAllLines(paths []string) domain.LinesInEachPath {
	jobs := make(chan string, len(paths))
	results := make(chan struct {
		path  string
		count int64
	}, len(paths))

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
				results <- struct {
					path  string
					count int64
				}{
					path:  path,
					count: n,
				}
			}
		}()
	}

	go func() {
		for _, p := range paths {
			jobs <- p
		}
		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	lineMap := make(domain.LinesInEachPath, len(paths))
	for r := range results {
		lineMap[r.path] = r.count
	}

	return lineMap
}
