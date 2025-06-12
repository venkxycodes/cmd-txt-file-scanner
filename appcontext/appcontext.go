package appcontext

import (
	"cmd-txt-file-scanner/aggregator"
	"cmd-txt-file-scanner/scanner"
	"cmd-txt-file-scanner/word-counter"
)

var appCtx *appContext

type appContext struct {
	fileScanner *scanner.FileScanner
	worker      *word_counter.WordCounter
	aggregator  *aggregator.Aggregator
}

func Init() {
	appCtx = &appContext{}
	appCtx.fileScanner = scanner.NewFileScanner()
	appCtx.worker = word_counter.NewWordCounter()
	appCtx.aggregator = aggregator.NewAggregator(*appCtx.worker)
}

func GetFileScanner() *scanner.FileScanner {
	return appCtx.fileScanner
}
func GetWorker() *word_counter.WordCounter {
	return appCtx.worker
}
func GetAggregator() *aggregator.Aggregator {
	return appCtx.aggregator
}
