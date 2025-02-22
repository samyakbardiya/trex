package ui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

func (m MainModel) View() string {
	var s string
	model := m.currentFocusedModel()
	if m.state == TextInput {
		s += lipgloss.JoinVertical(lipgloss.Top, focusedModelStyle.Render(m.textInput.View()), modelStyle.Render(m.viewport.View()))
	} else {
		s += lipgloss.JoinVertical(lipgloss.Top, modelStyle.Render(m.textInput.View()), focusedModelStyle.Render(m.viewport.View()))
	}
	s += helpStyle.Render(fmt.Sprintf("\ntab: focus next • n: new %s • q: exit\n", model))
	return s
}
