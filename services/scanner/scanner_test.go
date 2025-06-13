package scanner

import (
	"cmd-txt-file-scanner/utils"
	"strings"
	"testing"
)

func TestScanDirectory_ReturnsExpectedTxtFiles(t *testing.T) {
	tests := []struct {
		name             string
		setupFunc        func(t *testing.T) string
		expectedMinCount int
		description      string
	}{
		{
			name: "finds_txt_files_in_root_and_nested",
			setupFunc: func(t *testing.T) string {
				tempDir := t.TempDir()
				utils.CreateTestStructure(t, tempDir, map[string]interface{}{
					"a.txt":     "1",
					"b.TXT":     "2",
					"ignore.md": "skip",
					"nested": map[string]interface{}{
						"c.txt": "3",
						"d.log": "not included",
					},
				})
				return tempDir
			},
			expectedMinCount: 3,
		},
		{
			name: "returns_empty_slice_if_no_txt",
			setupFunc: func(t *testing.T) string {
				tempDir := t.TempDir()
				utils.CreateTestStructure(t, tempDir, map[string]interface{}{
					"notebook.pdf": "nope",
					"folder": map[string]interface{}{
						"img.png": "skip",
					},
				})
				return tempDir
			},
			expectedMinCount: 0,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			root := tc.setupFunc(t)
			fSc := NewFileScanner()
			files, err := fSc.Scan(root)

			if err != nil {
				t.Fatalf("Unexpected error: %v", err)
			}
			if len(files) < tc.expectedMinCount {
				t.Errorf("Expected at least %d .txt files, got %d", tc.expectedMinCount, len(files))
			}

			for _, f := range files {
				if !strings.HasSuffix(strings.ToLower(f), ".txt") {
					t.Errorf("Unexpected non-.txt file in results: %s", f)
				}
			}
		})
	}
}
