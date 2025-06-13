package domain

type TextFiles struct {
	Paths            []string
	NumberOfFiles    int64
	WordCountPerFile map[string]int64
}

type WordCount struct {
	Word  string
	Count int64
}

type MinHeap []WordCount

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i].Count < h[j].Count }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(WordCount)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[0 : n-1]
	return item
}
