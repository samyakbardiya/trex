package util

import (
	"fmt"
	"log"
	"regexp"
	"strings"
)

type MatchResult struct {
	InputText   string
	Highlighted string
	Pattern     string
	Matches     [][]int
}

func (mr *MatchResult) FindMatches() error {
	if len(mr.InputText) == 0 {
		return fmt.Errorf("empty text")
	}

	// TODO: implment customizable flags
	re, err := regexp.Compile("(?m)" + mr.Pattern)
	if err != nil {
		return fmt.Errorf("invalid regular expression: %q: %w", mr.Pattern, err)
	}

	mr.Matches = re.FindAllIndex([]byte(mr.InputText), -1)

	log.Printf("MATCHES:\n\texpr: %q\n\tmatches: %v", mr.Pattern, mr.Matches)
	return nil
}

func (mr *MatchResult) IsValidMatch(index int) bool {
	if index < 0 || index >= len(mr.Matches) {
		return false
	}

	match := mr.Matches[index]
	return len(match) == 2 &&
		match[0] >= 0 &&
		match[1] > match[0] &&
		match[1] <= len(mr.InputText)
}

func (mr *MatchResult) HighlightMatches(styleFunc func(...string) string) {
	if len(mr.InputText) == 0 || len(mr.Matches) == 0 {
		mr.Highlighted = mr.InputText
		return
	}

	var sb strings.Builder
	lastIndex := 0

	for i, match := range mr.Matches {
		if !mr.IsValidMatch(i) {
			continue
		}

		if match[0] > lastIndex {
			sb.WriteString(mr.InputText[lastIndex:match[0]])
		}

		matchedText := mr.InputText[match[0]:match[1]]
		sb.WriteString(styleFunc(matchedText))

		lastIndex = match[1]
	}

	if lastIndex < len(mr.InputText) {
		sb.WriteString(mr.InputText[lastIndex:])
	}

	mr.Highlighted = sb.String()
	log.Printf("highlighted:\n%q\n%s", mr.Highlighted, mr.Highlighted)
}
