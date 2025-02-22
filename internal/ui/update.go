package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
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
	m.viewport.SetContent(m.matchRes.Highlighted)
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

	if m.matchRes.Pattern != m.input.Value() {
		m.matchRes.Pattern = m.input.Value()
		m.updateRegexMatches()
	}

	return m, cmd
}

func (m *model) updateRegexMatches() {
	if err := m.matchRes.FindMatches(); err != nil {
		m.err = err
		return
	}
	m.matchRes.HighlightMatches(tsHighlight)
	m.viewport.SetContent(m.matchRes.Highlighted)
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
