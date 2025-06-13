package counter

import (
	"cmd-txt-file-scanner/utils"
	"path/filepath"
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []string
	}{
		{
			input:    "Hello World",
			expected: []string{"hello", "world"},
		},
		{
			input:    " 123 Go!! Gophers",
			expected: []string{"go", "gophers"},
		},
		{
			input:    "No-symbols-here",
			expected: []string{"no", "symbols", "here"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			"$$$ %%%",
			[]string{},
		},
	}
	for _, tt := range tests {
		got := tokenize(tt.input)
		if !reflect.DeepEqual(got, tt.expected) {
			t.Errorf("tokenize(%q) = %v; want %v", tt.input, got, tt.expected)
		}
	}
}

func TestCountWords(t *testing.T) {
	tests := []struct {
		name      string
		structure map[string]interface{}
		expected  map[string]int64
		filename  string
	}{
		{
			name: "simple_count",
			structure: map[string]interface{}{
				"file.txt": "Hello world\nhello again",
			},
			expected: map[string]int64{
				"hello": 2,
				"world": 1,
				"again": 1,
			},
			filename: "file.txt",
		},
		{
			name: "with_symbols",
			structure: map[string]interface{}{
				"sym.txt": "Go! Go? Go.\n123",
			},
			expected: map[string]int64{
				"go": 3,
			},
			filename: "sym.txt",
		},
		{
			name: "empty_file",
			structure: map[string]interface{}{
				"empty.txt": "",
			},
			expected: map[string]int64{},
			filename: "empty.txt",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tempDir := t.TempDir()
			utils.CreateTestStructure(t, tempDir, tt.structure)

			fullPath := filepath.Join(tempDir, tt.filename)
			w := &WordCounter{}

			result, err := w.CountWords(fullPath)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("CountWords(%s) = %v; want %v", tt.filename, result, tt.expected)
			}
		})
	}
}
