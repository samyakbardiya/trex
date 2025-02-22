package util

import (
	"fmt"
	"os"
	"path/filepath"
)

// ValidateFilepath returns the absolute file path for the provided path if it points to an existing file. It cleans the input, resolves it to an absolute path, and verifies that the path exists and is not a directory. If the resolution, existence check, or file type validation fails, an error describing the issue is returned.
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
