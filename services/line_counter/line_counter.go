package line_counter

import (
	"bufio"
	"os"
)

type LineCounter struct{}

func NewLineCounter() *LineCounter {
	return &LineCounter{}
}

func (lc *LineCounter) CountLines(path string) (int64, error) {
	f, err := os.Open(path)
	if err != nil {
		return 0, err
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	lines := 0
	for scanner.Scan() {
		lines++
	}
	return int64(lines), scanner.Err()
}
