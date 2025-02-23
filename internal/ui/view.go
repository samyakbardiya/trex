package ui

import (
	"fmt"
	"log"
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
	case stateNotification:
		return m.renderBox("RegEx copied to clipboard!", bsSuccess)
	case stateExiting:
		return m.renderBox("Do you really want to quit? [y/N]", bsError)
	default:
		return ""
	}
}

func (m model) renderInputField() string {
	inputText := tsNormal.Render("> ") + tsHelp.Render("/") + m.input.View() + tsHelp.Render("/gm")
	if m.err != nil {
		log.Print("error:", m.err)
		return bsError.Width(m.width - 4).Render(inputText)
	}
	if m.focus == focusInput {
		return bsFocus.Width(m.width - 4).Render(inputText)
	}
	return bsUnfocus.Width(m.width - 4).Render(inputText)
}

func (m model) renderContentView() string {
	if m.focus == focusContent {
		return bsFocus.Render(m.viewport.View())
	}
	return bsUnfocus.Render(m.viewport.View())
}

func (m model) renderCentered(str string) string {
	return lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, str)
}

func (m model) renderBox(message string, style lipgloss.Style) string {
	box := style.Padding(1, 4).Render(message)
	if w, _ := lipgloss.Size(box); m.width < w {
		return m.renderCentered(lipgloss.NewStyle().Width(m.width).Render(message))
	}
	return m.renderCentered(box)
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
