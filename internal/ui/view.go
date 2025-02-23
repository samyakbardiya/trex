package ui

import (
	"github.com/charmbracelet/lipgloss"
)

const (
	helpText = "\ntab: focus next • q: exit\n"
)

func (m model) View() string {
	switch m.state {
	case stateActive:
		s := []string{m.renderInputField(), m.renderContentView(), m.renderHelpText()}
		return lipgloss.JoinVertical(lipgloss.Top, s...)
	case stateAlertClipboard:
		return m.renderSuccessBox("RegEx copied to clipboard!")
	case stateExiting:
		return m.renderErrorBox("Do you really want to quit? [y/N]")
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

func (m model) renderErrorBox(message string) string {
	box := bsError.Padding(1, 4).Render(message)
	if w, h := lipgloss.Size(box); m.width < w || m.height < h {
		return lipgloss.NewStyle().Width(m.width).Render(message)
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}

func (m model) renderSuccessBox(message string) string {
	box := bsSuccess.Padding(1, 4).Render(message)
	if w, h := lipgloss.Size(box); m.width < w || m.height < h {
		return lipgloss.NewStyle().Width(m.width).Render(message)
	}
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, box)
}
