package ui

import (
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

func (m MainModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "tab":
			if m.state == TextInput {
				m.state = ContentView
			} else {
				m.state = TextInput
			}
		}
		switch m.state {
		// update whichever model is focused
		case ContentView:
			m.viewport, cmd = m.viewport.Update(msg)
			cmds = append(cmds, cmd)
		default:
			m.textInput, cmd = m.textInput.Update(msg)
			cmds = append(cmds, cmd)
		}

		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}
	case ErrMsg:
		m.err = msg
		return m, nil

	case tea.WindowSizeMsg:
		customWidth := msg.Width - 2
		customHeight := msg.Height - 15

		m.viewport = viewport.New(customWidth, customHeight)
		m.viewport.SetContent(m.content)
	}

	return m, tea.Batch(cmds...)
}
