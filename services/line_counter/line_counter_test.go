package line_counter

import (
	"cmd-txt-file-scanner/utils"
	"path/filepath"
	"testing"
)

func TestCountLines(t *testing.T) {
	tests := []struct {
		name      string
		structure map[string]interface{}
		filename  string
		want      int64
	}{
		{
			name:      "test empty file",
			structure: map[string]interface{}{"empty.txt": ""},
			filename:  "empty.txt",
			want:      0,
		},
		{
			name:      "test single line no newline",
			structure: map[string]interface{}{"single.txt": "hello world"},
			filename:  "single.txt",
			want:      1,
		},
		{
			name:      "test single line with newline",
			structure: map[string]interface{}{"single_nl.txt": "hello world\n"},
			filename:  "single_nl.txt",
			want:      1,
		},
		{
			name:      "test multiple lines",
			structure: map[string]interface{}{"dir1": map[string]interface{}{"multi.txt": "line1\nline2\nline3\n"}},
			filename:  filepath.Join("dir1", "multi.txt"),
			want:      3,
		},
		{
			name:      "test trailing blank lines",
			structure: map[string]interface{}{"dir4": map[string]interface{}{"blank.txt": "a\nb\n\n"}},
			filename:  filepath.Join("dir4", "blank.txt"),
			want:      3,
		},
	}
	lc := NewLineCounter()
	tempDir := t.TempDir()

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			utils.CreateTestStructure(t, tempDir, tc.structure)
			fullPath := filepath.Join(tempDir, tc.filename)

			got, err := lc.CountLines(fullPath)
			if err != nil {
				t.Fatalf("CountLines(%s) returned error: %v", tc.filename, err)
			}
			if got != tc.want {
				t.Errorf("CountLines(%s) = %d; want %d", tc.filename, got, tc.want)
			}
		})
	}

	t.Run("nonexistent file", func(t *testing.T) {
		if _, err := lc.CountLines(filepath.Join(tempDir, "no", "file.txt")); err == nil {
			t.Error("expected error for nonexistent file, got nil")
		}
	})
}
