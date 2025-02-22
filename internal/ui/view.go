package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m Model) View() string {
	var s string
	switch m.state {
	case TextInput:
		s += lipgloss.JoinVertical(lipgloss.Top, focusedModelStyle.Render(m.textInput.View()), modelStyle.Render(m.viewport.View()))
	case ContentView:
		s += lipgloss.JoinVertical(lipgloss.Top, modelStyle.Render(m.textInput.View()), focusedModelStyle.Render(m.viewport.View()))
	}
	currentFocus := m.currentFocusedModel()
	s += helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", currentFocus))
	return s
}
