package appcontext

import (
	"cmd-txt-file-scanner/services/scanner"
	"cmd-txt-file-scanner/services/word-counter"
)

var appCtx *appContext

type appContext struct {
	fileScanner *scanner.FileScanner
	wordCounter *word_counter.WordCounter
}

func Init() {
	appCtx = &appContext{}
	appCtx.fileScanner = scanner.NewFileScanner()
	appCtx.wordCounter = word_counter.NewWordCounter()
}

func GetFileScanner() *scanner.FileScanner {
	return appCtx.fileScanner
}
func GetWordCounter() *word_counter.WordCounter {
	return appCtx.wordCounter
}
