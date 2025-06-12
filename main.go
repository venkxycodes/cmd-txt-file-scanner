package main

import (
	"cmd-txt-file-scanner/domain"
	"cmd-txt-file-scanner/scanner"
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <directory>", os.Args[0])
	}
	rootDir := os.Args[1]

	filePaths, err := scanner.ScanDirectory(rootDir)
	textFiles := domain.TextFiles{
		Paths: filePaths,
		Count: int64(len(filePaths)),
	}
	if err != nil {
		log.Fatalf("err-scanning-directory: %v", err)
	}
	for _, path := range textFiles.Paths {
		fmt.Println(path)
	}
	fmt.Println(textFiles.Count)
}
