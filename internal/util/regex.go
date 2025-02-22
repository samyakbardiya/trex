package util

import (
	"fmt"
	"regexp"
)

func FindMatches(expr string, text []byte) ([][]int, error) {
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
	if matches == nil {
		matches = [][]int{}
	}

	return matches, nil
}

func IsValidMatch(match []int, contentLen int) bool {
	return len(match) == 2 &&
		match[0] >= 0 &&
		match[1] > match[0] &&
		match[1] <= contentLen
}
