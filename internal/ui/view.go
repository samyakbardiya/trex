package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	helpText = "\ntab: focus next • q: exit\n"
)

func (m model) View() string {
	switch m.state {
	case stateWorking:
		s := []string{m.renderInputField(), m.renderContentView(), m.renderHelpText()}
		return lipgloss.JoinVertical(lipgloss.Top, s...)
	case stateQuiting:
		return m.renderQuitBox()
	default:
		return ""
	}
}

func (m model) renderInputField() string {
	if m.err != nil {
		return bsError.Render(m.input.View())
	}
	if m.focus == focusInput {
		return bsFocus.Render(m.input.View())
	}
	return bsUnfocus.Render(m.input.View())
}

func (m model) renderContentView() string {
	if m.focus == focusContent {
		return bsFocus.Render(m.viewport.View())
	}
	return bsUnfocus.Render(m.viewport.View())
}

func (m model) renderHelpText() string {
	return tsHelp.Render(helpText)
}

func (m model) renderQuitBox() string {
	question := "Do you really want to quit? [y/N]"
	dialog := bsError.Padding(1, 4).Render(question)
	if w, h := lipgloss.Size(dialog); m.width < w || m.height < h {
		return question
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, dialog)
}
