package ui

import (
	"fmt"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {

	// TODO: Handle Error
	case ErrMsg:
		m.err = msg
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			// update whichever model is focused
			switch m.state {
			case TextInput:
				m.state = ContentView
			case ContentView:
				m.state = TextInput
			default:
				m.err = fmt.Errorf("invalid state")
			}
		}

		// pass KeyMsg to the focused-box
		switch m.state {
		case TextInput:
			m.textInput, cmd = m.textInput.Update(msg)
			cmds = append(cmds, cmd)
		case ContentView:
			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
		}

	case tea.WindowSizeMsg:
		customWidth := msg.Width - 2
		customHeight := msg.Height - 15

		m.viewport = viewport.New(customWidth, customHeight)
		m.viewport.SetContent(m.content)
	}

	return m, tea.Batch(cmds...)
}
