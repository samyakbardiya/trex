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
	if m.width < minWidth || m.height < minHeight {
		return m.renderBox("Window too small!\nPlease resize.", bsError)
	}

	switch m.state {
	case stateActive:
		return lipgloss.JoinVertical(
			lipgloss.Center,
			lipgloss.JoinHorizontal(
				lipgloss.Center,
				lipgloss.JoinVertical(
					lipgloss.Center,
					m.renderInputField(),
					m.renderContentView(),
				),
				m.renderCheatsheet(),
			),
			m.renderHelpText(),
		)
	case stateNotification:
		return m.renderBox("RegEx copied to clipboard!", bsSuccess)
	case stateExiting:
		return m.renderBox("Do you really want to quit? [y/N]", bsError)
	default:
		return ""
	}
}

func (m model) renderInputField() string {
	text := tsNormal.Render("TReX => ") + tsHelp.Render("/") + m.input.View() + tsHelp.Render("/gm")
	style := bsUnfocus
	if m.err != nil {
		style = bsError
	} else if m.focus == focusInput {
		style = bsFocus
	}
	maxWidth := int(float32(m.width)*leftWidthRatio) - borderWidthDiff
	return style.Width(maxWidth).Render(text)
}

func (m model) renderContentView() string {
	if m.focus == focusContent {
		return bsFocus.Render(m.viewport.View())
	}
	return bsUnfocus.Render(m.viewport.View())
}

func (m model) renderCheatsheet() string {
	view := m.cheatsheet.View()
	if m.focus == focusCheatsheet {
		return bsFocus.Render(view)
	}
	return bsUnfocus.Render(view)
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
			{description: "clear", binding: tea.KeyCtrlW.String()},
			{description: "copy regex", binding: tea.KeyEnter.String()},
		}
	case focusContent:
		focusSpecificKeyBindings = []keyBinding{
			{description: "scroll up", binding: tea.KeyUp.String()},
			{description: "scroll down", binding: tea.KeyDown.String()},
		}
	case focusCheatsheet:
		focusSpecificKeyBindings = []keyBinding{
			{description: "up", binding: tea.KeyUp.String()},
			{description: "down", binding: tea.KeyDown.String()},
			{description: "next page", binding: tea.KeyRight.String()},
			{description: "prev page", binding: tea.KeyLeft.String()},
			{description: "goto start", binding: tea.KeyHome.String()},
			{description: "goto end", binding: tea.KeyEnd.String()},
		}
	}

	baseKeyBindings := []keyBinding{
		{description: "Focus Next", binding: tea.KeyTab.String()},
		{description: "Quit", binding: tea.KeyEsc.String()},
		{description: "Force Quit", binding: tea.KeyCtrlC.String()},
	}
	keyBinding := append(focusSpecificKeyBindings, baseKeyBindings...)

	var keyMap []string
	for _, kb := range keyBinding {
		keyMap = append(keyMap, fmt.Sprintf("%s: <%s>", kb.description, kb.binding))
	}
	return tsHelp.MaxWidth(m.width).Render(strings.Join(keyMap, " | "))
}
