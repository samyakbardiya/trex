package util

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func GetFilePath(path string) (string, error) {
	absPath, err := resolveFilePath(path)
	if err != nil {
		return "", err
	}

	if err := validateFile(absPath); err != nil {
		return "", err
	}

	log.Println("filepath:", absPath)
	return absPath, nil
}

func resolveFilePath(path string) (string, error) {
	absPath, err := filepath.Abs(filepath.Clean(path))
	if err != nil {
		return "", fmt.Errorf("failed to resolve path %q: %w", path, err)
	}

	if realPath, err := filepath.EvalSymlinks(absPath); err == nil {
		absPath = realPath
	}

	return absPath, nil
}

func validateFile(path string) error {
	const maxFileSize = 10 * 1024 * 1024 // 10MB limit

	fileInfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file does not exist: %s", path)
		}
		return fmt.Errorf("failed to check file %q: %w", path, err)
	}

	if fileInfo.Size() > maxFileSize {
		return fmt.Errorf("file size exceeds limit of %d bytes", maxFileSize)
	}

	if fileInfo.IsDir() {
		return fmt.Errorf("path is a directory, not a file: %s", path)
	}

	return nil
}
