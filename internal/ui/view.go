package ui

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
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

func (m model) renderHelpText() string {
	var focusSpecificKeyBindings []keyBinding
	switch m.focus {
	case focusInput:
		focusSpecificKeyBindings = []keyBinding{
			{description: "Clear", binding: tea.KeyCtrlW.String()},
			{description: "Copy RegEx", binding: tea.KeyEnter.String()},
		}
	case focusContent:
		focusSpecificKeyBindings = []keyBinding{
			{description: "Scroll Up", binding: tea.KeyUp.String()},
			{description: "Scroll Down", binding: tea.KeyDown.String()},
		}
	}

	baseKeyBindings := []keyBinding{
		{description: "Focus Next", binding: tea.KeyTab.String()},
		{description: "Quit", binding: "q"},
		{description: "Force Quit", binding: tea.KeyCtrlC.String()},
	}
	keyBinding := append(focusSpecificKeyBindings, baseKeyBindings...)

	var sb strings.Builder
	for i, kb := range keyBinding {
		fmt.Fprintf(&sb, "%s: <%s>", kb.description, kb.binding)
		if i < len(keyBinding)-1 {
			sb.WriteString(" | ")
		}
	}

	return tsHelp.Render(sb.String())
}
