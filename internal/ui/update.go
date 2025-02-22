package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

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
	case tea.KeyEsc:
		return m.toggleState()
	case tea.KeyTab:
		return m.getNextFocus()
	case tea.KeyCtrlC:
		return m, tea.Quit
	}

	if m.state == stateQuiting {
		switch msg.String() {
		case "y", "Y":
			return m, tea.Quit
		default:
			return m.toggleState()
		}
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
	m.width = msg.Width
	m.height = msg.Height
	m.viewport = viewport.New(m.width-4, m.height-8)
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
	if m.err = m.matchRes.FindMatches(); m.err != nil {
		return
	}
	m.matchRes.HighlightMatches(tsHighlight.Render)
	m.viewport.SetContent(m.matchRes.Highlighted)
}

func (m *model) getNextFocus() (tea.Model, tea.Cmd) {
	switch m.focus {
	case focusInput:
		m.focus = focusContent
	case focusContent:
		m.focus = focusInput
	default:
		m.focus = focusInput
	}
	return m, nil
}

func (m *model) toggleState() (tea.Model, tea.Cmd) {
	switch m.state {
	case stateWorking:
		m.state = stateQuiting
	case stateQuiting:
		m.state = stateWorking
	default:
		m.state = stateWorking
	}
	return m, nil
}
