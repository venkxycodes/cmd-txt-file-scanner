package main

import (
	"cmd-txt-file-scanner/appcontext"
	"cmd-txt-file-scanner/domain"
	"fmt"
	"log"
	"os"
)

func execute(rootDir string) {
	paths, err := appcontext.GetFileScanner().Scan(rootDir)
	if err != nil {
		log.Fatalf("Failed to scan directory: %v", err)
		return
	}
	directoryDetails := domain.TextFiles{
		Paths:         paths,
		NumberOfFiles: int64(len(paths)),
	}
	wordCounts, err := appcontext.GetAggregator().AggregateWordCounts(paths)
	if err != nil {
		log.Fatalf("Failed to aggregate word counts: %v", err)
	}

	for word, count := range wordCounts {
		fmt.Printf("%s: %d\n", word, count)
	}
	fmt.Printf("Processed %d files\n", len(paths))
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <directory>", os.Args[0])
	}
	rootDir := os.Args[1]
	appcontext.Init()
	execute(rootDir)
}
