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
		return m.renderBox("RegEx copied to clipboard!", bsSuccess)
	case stateExiting:
		return m.renderBox("Do you really want to quit? [y/N]", bsError)
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

func (m model) renderBox(message string, style lipgloss.Style) string {
	box := style.Padding(1, 4).Render(message)
	if w, _ := lipgloss.Size(box); m.width < w {
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

	var keyMap []string
	for _, kb := range keyBinding {
		keyMap = append(keyMap, fmt.Sprintf("%s: <%s>", kb.description, kb.binding))
	}
	return tsHelp.Render(strings.Join(keyMap, " | "))
}
