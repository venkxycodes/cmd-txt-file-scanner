package appcontext

import (
	"cmd-txt-file-scanner/services/line_counter"
	"cmd-txt-file-scanner/services/scanner"
	"cmd-txt-file-scanner/services/word_counter"
)

var appCtx *appContext

type appContext struct {
	fileScanner *scanner.FileScanner
	wordCounter *word_counter.WordCounter
	lineCounter *line_counter.LineCounter
}

func Init() {
	appCtx = &appContext{}
	appCtx.fileScanner = scanner.NewFileScanner()
	appCtx.wordCounter = word_counter.NewWordCounter()
	appCtx.lineCounter = line_counter.NewLineCounter()
}

func GetFileScanner() *scanner.FileScanner {
	return appCtx.fileScanner
}
func GetWordCounter() *word_counter.WordCounter {
	return appCtx.wordCounter
}
func GetLineCounter() *line_counter.LineCounter {
	return appCtx.lineCounter
}
