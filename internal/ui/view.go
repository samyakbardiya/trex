package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	helpText = "\ntab: focus next • q: exit\n"
)

func (m model) View() string {
	return lipgloss.JoinVertical(
		lipgloss.Top,
		m.renderInputField(),
		m.renderContentView(),
		m.renderHelpText(),
	)
}

func (m model) renderInputField() string {
	if m.err != nil {
		return bsError(m.input.View())
	}
	if m.focus == focusInput {
		return bsFocus(m.input.View())
	}
	return bsUnfocus(m.input.View())
}

func (m model) renderContentView() string {
	if m.focus == focusContent {
		return bsFocus(m.viewport.View())
	}
	return bsUnfocus(m.viewport.View())
}

func (m model) renderHelpText() string {
	return tsHelp(helpText)
}
