package main

import (
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
	if err != nil {
		log.Fatalf("err-scanning-directory: %v", err)
	}
	fmt.Println(filePaths)
}
