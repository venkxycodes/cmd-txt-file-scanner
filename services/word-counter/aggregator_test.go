package counter

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
			name: "aggregates_words_across_multiple_files",
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
			name: "includes_empty_file_in_aggregation",
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
			name:      "returns_empty_map_on_no_files",
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
			got, err := wc.AggregateWordCounts(fullPaths)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("AggregateWordCounts() = %v; want %v", got, tc.expected)
			}
		})
	}
}
