package ui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samyakbardiya/trex/internal/util"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	// TODO: Handle Error
	case ErrMsg:
		m.err = msg
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			// update whichever model is focused
			switch m.state {
			case TextInput:
				m.state = ContentView
			case ContentView:
				m.state = TextInput
			default:
				m.err = fmt.Errorf("invalid state")
			}
		}

		// pass KeyMsg to the focused-box
		switch m.state {
		case TextInput:
			m.textInput, cmd = m.textInput.Update(msg)
			m.expr = m.textInput.Value()
			m.matches, m.err = util.FindMatches(m.expr, []byte(m.content))
			m.highlighted = highlightContent(m.content, m.matches, tsHighlight)
			m.viewport.SetContent(m.highlighted)
			cmds = append(cmds, cmd)
		case ContentView:
			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
		}

	case tea.WindowSizeMsg:
		customWidth := msg.Width - 2
		customHeight := msg.Height - 15

		m.viewport = viewport.New(customWidth, customHeight)
		m.viewport.SetContent(m.content)
	}

	return m, tea.Batch(cmds...)
}

func highlightContent(
	content string,
	matches [][]int,
	styleFunc func(...string) string,
) string {
	contentLen := len(content)
	if len(matches) == 0 || contentLen == 0 {
		return content
	}

	var sb strings.Builder
	sb.Grow(len(content) * 2)

	for i := len(matches) - 1; i >= 0; i-- {
		match := matches[i]
		if !util.IsValidMatch(match, contentLen) {
			continue
		}

		matchedText := content[match[0]:match[1]]
		styledText := styleFunc(matchedText)

		sb.WriteString(content[:match[0]])
		sb.WriteString(styledText)
		sb.WriteString(content[match[1]:])

		content = sb.String()
		sb.Reset()
		sb.Grow(len(content) * 2)
	}

	return content
}
