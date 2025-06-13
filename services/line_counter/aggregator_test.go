package line_counter

import (
	"cmd-txt-file-scanner/utils"
	"path/filepath"
	"reflect"
	"testing"
)

func TestCountAllLines(t *testing.T) {
	tests := []struct {
		name      string
		structure map[string]interface{}
		expected  map[string]int64
		files     []string
	}{
		{
			name: "count lines in multiple files",
			structure: map[string]interface{}{
				"a.txt": "Hello world\nAnother line",
				"b.txt": "Line one\nLine two\nLine three",
				"c.txt": "Single line",
			},
			expected: map[string]int64{
				"a.txt": 2,
				"b.txt": 3,
				"c.txt": 1,
			},
			files: []string{"a.txt", "b.txt", "c.txt"},
		},
		{
			name: "includes empty file",
			structure: map[string]interface{}{
				"one.txt":   "Just one line",
				"empty.txt": "",
				"two.txt":   "Line one\nLine two",
			},
			expected: map[string]int64{
				"one.txt":   1,
				"empty.txt": 0,
				"two.txt":   2,
			},
			files: []string{"one.txt", "empty.txt", "two.txt"},
		},
		{
			name:      "no files",
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

			rawResult := CountAllLines(fullPaths)

			// Normalize keys to base filenames for comparison
			got := make(map[string]int64)
			for fullPath, count := range rawResult {
				got[filepath.Base(fullPath)] = count
			}

			if !reflect.DeepEqual(got, tc.expected) {
				t.Errorf("CountAllLines() = %v; want %v", got, tc.expected)
			}
		})
	}
}
