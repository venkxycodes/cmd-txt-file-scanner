package domain

type TextFiles struct {
	Paths            []string
	NumberOfFiles    int64
	WordCountPerFile map[string]int64
}
