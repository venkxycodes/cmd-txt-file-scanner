package word_counter

import (
	"cmd-txt-file-scanner/utils"
	"path/filepath"
	"reflect"
	"testing"
)

func TestAggregateWordCounts(t *testing.T) {
	tests := []struct {
		name      string
		structure map[string]interface{}
		expected  map[string]int64
		files     []string
	}{
		{
			name: "test aggregates words across multiple files",
			structure: map[string]interface{}{
				"a.txt": "Hello world",
				"b.txt": "Hello Go",
				"c.txt": "Go Go Go",
			},
			expected: map[string]int64{
				"hello": 2,
				"world": 1,
				"go":    4,
			},
			files: []string{"a.txt", "b.txt", "c.txt"},
		},
		{
			name: "test includes empty file in aggregation",
			structure: map[string]interface{}{
				"one.txt":   "Hello there",
				"empty.txt": "",
				"two.txt":   "there again",
			},
			expected: map[string]int64{
				"hello": 1,
				"there": 2,
				"again": 1,
			},
			files: []string{"one.txt", "empty.txt", "two.txt"},
		},
		{
			name:      "test returns empty map on no files",
			structure: map[string]interface{}{},
			expected:  map[string]int64{},
			files:     []string{},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			tempDir := t.TempDir()
			utils.CreateTestStructure(t, tempDir, tc.structure)

			var fullPaths []string
			for _, f := range tc.files {
				fullPaths = append(fullPaths, filepath.Join(tempDir, f))
			}

			wc := &WordCounter{}
			got, err := wc.CountWordsInAllFiles(fullPaths)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("CountWordsInAllFiles() = %v; want %v", got, tc.expected)
			}
		})
	}
}
