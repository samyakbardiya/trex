package util

import (
	"fmt"
	"regexp"
)

func GetAllMatchingIndex(expr string, text []byte) ([][]int, error) {
	if expr == "" {
		return nil, fmt.Errorf("empty regular expression")
	}
	if len(text) == 0 {
		return nil, fmt.Errorf("empty text")
	}

	re, err := regexp.Compile(expr)
	if err != nil {
		return nil, fmt.Errorf("invalid regular expression: %q: %w", expr, err)
	}

	matches := re.FindAllIndex(text, -1)
	return matches, nil
}
