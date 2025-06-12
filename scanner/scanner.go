package scanner

import (
	"os"
	"path/filepath"
	"strings"
)

func ScanDirectory(root string) ([]string, error) {
	var filePaths []string

	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".txt") {
			filePaths = append(filePaths, path)
		}
		return nil
	})

	return filePaths, err
}
