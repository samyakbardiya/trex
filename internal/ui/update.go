package ui

import (
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samyakbardiya/trex/internal/util"
)

// Update handles all UI state updates and user input
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg)
	case tea.MouseMsg:
		return m.handleMouseMsg(msg)
	default:
		return m, nil
	}
}

func (m model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEsc, tea.KeyCtrlC:
		return m, tea.Quit
	case tea.KeyTab:
		m.focus = m.getNextFocus()
	}

	var cmd tea.Cmd
	switch m.focus {
	case focusInput:
		return m.handleInputUpdate(msg)
	case focusContent:
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.viewport = viewport.New(msg.Width-4, msg.Height-8)
	m.viewport.SetContent(m.regexData.Highlighted)
	return m, nil
}

func (m model) handleMouseMsg(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) handleInputUpdate(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	if m.regexData.Regexpr != m.input.Value() {
		m.regexData.Regexpr = m.input.Value()
		m.updateRegexMatches()
	}

	return m, cmd
}

func (m *model) updateRegexMatches() {
	var err error
	m.regexData.Matches, err = util.FindMatches(
		m.regexData.Regexpr,
		[]byte(m.regexData.Raw),
	)
	if err != nil {
		m.err = err
		return
	}

	highlighted := highlightMatches(m.regexData.Raw, m.regexData.Matches)
	m.regexData.Highlighted = highlighted
	m.viewport.SetContent(highlighted)
}

func (m model) getNextFocus() focus {
	switch m.focus {
	case focusInput:
		return focusContent
	case focusContent:
		return focusInput
	default:
		return focusInput
	}
}

func highlightMatches(content string, matches [][]int) string {
	if len(matches) == 0 || len(content) == 0 {
		return content
	}

	segments := make([]string, 0, len(matches)*2+1)
	lastIndex := 0

	// Process matches in order
	for _, match := range matches {
		if !util.IsValidMatch(match, len(content)) {
			continue
		}

		// Add text before match
		if match[0] > lastIndex {
			segments = append(segments, content[lastIndex:match[0]])
		}

		// Add highlighted match
		matchedText := content[match[0]:match[1]]
		segments = append(segments, tsHighlight(matchedText))
		lastIndex = match[1]
	}

	// Add remaining text after last match
	if lastIndex < len(content) {
		segments = append(segments, content[lastIndex:])
	}

	return strings.Join(segments, "")
}
