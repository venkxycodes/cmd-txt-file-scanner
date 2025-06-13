package domain

type FileStat struct {
	Path      string
	LineCount int
}

type TextFiles struct {
	Paths            []string
	NumberOfFiles    int64
	WordCountPerFile map[string]int64
}

type WordCount struct {
	Word  string
	Count int64
}

type LinesInEachPath map[string]int64 // Key is path, value is count
