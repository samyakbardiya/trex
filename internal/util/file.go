package util

import (
	"fmt"
	"os"
	"path/filepath"
)

func ValidateFilepath(path string) (string, error) {
	absPath, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return "", fmt.Errorf("failed to resolve path %q: %w", path, err)
	}

	fileInfo, err := os.Stat(absPath)
	if err != nil {
		if os.IsNotExist(err) {
			return "", fmt.Errorf("file does not exists: %s", absPath)
		}
		return "", fmt.Errorf("failed to check file %q: %w", absPath, err)
	}

	if fileInfo.IsDir() {
		return "", fmt.Errorf("path is a directory, not a file: %s", absPath)
	}

	return absPath, nil
}
