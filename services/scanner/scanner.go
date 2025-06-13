package scanner

import (
	"io/fs"
	"path/filepath"
	"strings"
)

type FileScanner struct{}

func NewFileScanner() *FileScanner {
	return &FileScanner{}
}

func (s *FileScanner) Scan(root string) ([]string, error) {
	var filePaths []string
	err := filepath.WalkDir(root, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !d.IsDir() && strings.HasSuffix(strings.ToLower(d.Name()), ".txt") {
			filePaths = append(filePaths, path)
		}
		return nil
	})
	return filePaths, err
}
