package util

import (
	"fmt"
	"log"
	"regexp"
)

type RegexMatch struct {
	Raw         string
	Highlighted string
	Regexpr     string
	Matches     [][]int
}

func FindMatches(expr string, text []byte) ([][]int, error) {
	if len(text) == 0 {
		return nil, fmt.Errorf("empty text")
	}

	// TODO: implment customizable flags
	re, err := regexp.Compile("(?m)" + expr)
	if err != nil {
		return nil, fmt.Errorf("invalid regular expression: %q: %w", expr, err)
	}

	matches := re.FindAllIndex(text, -1)
	if matches == nil {
		matches = [][]int{}
	}

	log.Printf("MATCHES:\n\texpr: %q\n\tmatches: %v", expr, matches)
	return matches, nil
}

func IsValidMatch(match []int, contentLen int) bool {
	return len(match) == 2 &&
		match[0] >= 0 &&
		match[1] > match[0] &&
		match[1] <= contentLen
}
