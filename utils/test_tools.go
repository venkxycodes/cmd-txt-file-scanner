package utils

import (
	"os"
	"path/filepath"
	"testing"
)

func CreateTestStructure(t *testing.T, baseDir string, structure map[string]interface{}) {
	t.Helper()
	for name, value := range structure {
		fullPath := filepath.Join(baseDir, name)
		switch v := value.(type) {
		case string:
			if err := os.WriteFile(fullPath, []byte(v), 0644); err != nil {
				t.Fatalf("Failed to create file %s: %v", fullPath, err)
			}
		case map[string]interface{}:
			if err := os.Mkdir(fullPath, 0755); err != nil {
				t.Fatalf("Failed to create dir %s: %v", fullPath, err)
			}
			CreateTestStructure(t, fullPath, v)
		default:
			t.Fatalf("Unsupported type for test structure: %T", v)
		}
	}
}
