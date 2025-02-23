package ui

import (
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"golang.design/x/clipboard"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		return m.handleKeyMsg(msg)
	case tea.WindowSizeMsg:
		return m.handleWindowSizeMsg(msg)
	case tea.MouseMsg:
		return m.handleMouseMsg(msg)
	case tickMsg:
		return m.handleTickMsg()
	default:
		return m, nil
	}
}

func (m model) handleKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyEsc:
		return m.toggleState()
	case tea.KeyTab:
		return m.getNextFocus()
	case tea.KeyCtrlC:
		return m, tea.Quit
	}

	switch m.state {
	case stateExiting:
		switch msg.String() {
		case "y", "Y":
			return m, tea.Quit
		default:
			return m.toggleState()
		}
	case stateNotification:
		return m, nil // blocks KeyMsg
	}

	var cmd tea.Cmd
	switch m.focus {
	case focusInput:
		return m.handleInputUpdate(msg)
	case focusContent:
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	case focusCheatsheet:
		return m.handleCheatsheetKeyMsg(msg)
	}

	return m, nil
}

func (m model) handleWindowSizeMsg(msg tea.WindowSizeMsg) (tea.Model, tea.Cmd) {
	m.width = msg.Width
	m.height = msg.Height

	maxViewportWidth := int(float32(m.width)*leftWidthRatio) - borderWidthDiff
	m.viewport.Width = maxViewportWidth
	m.viewport.Height = m.height - minHelpHeight - minInputHeight

	maxCheatsheetWidth := int(float32(m.width)*rightWidthRatio) - borderWidthDiff
	m.cheatsheet.SetWidth(maxCheatsheetWidth)
	m.cheatsheet.SetHeight(m.height - minHelpHeight - 1)

	return m, nil
}

func (m model) handleMouseMsg(msg tea.MouseMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	m.cheatsheet, cmd = m.cheatsheet.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) handleTickMsg() (tea.Model, tea.Cmd) {
	switch m.state {
	case stateNotification:
		m.state = stateActive // resets state
	}
	return m, nil
}

func (m *model) handleInputUpdate(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.input, cmd = m.input.Update(msg)

	switch msg.Type {
	case tea.KeyEnter:
		if err := clipboard.Init(); err == nil {
			clipboard.Write(clipboard.FmtText, []byte(m.matchRes.Pattern))
			m.state = stateNotification
			return m, tea.Batch(cmd, timeout(1*time.Second))
		}
	default:
		if m.matchRes.Pattern != m.input.Value() {
			m.matchRes.Pattern = m.input.Value()
			m.updateRegexMatches()
		}
	}

	return m, cmd
}

func (m *model) handleCheatsheetKeyMsg(msg tea.KeyMsg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.cheatsheet, cmd = m.cheatsheet.Update(msg)

	switch msg.Type {
	case tea.KeyEnter:
		i := m.cheatsheet.Index()
		items := m.cheatsheet.Items()
		inValue := m.input.Value()

		if selected, ok := items[i].(item); ok {
			switch selected.itemType {
			case itemCheatcode:
				m.matchRes.Pattern = inValue + selected.pattern
				m.input.SetValue(m.matchRes.Pattern)
			case itemTemplate:
				m.matchRes.Pattern = selected.pattern
				m.matchRes.InputText = selected.testStr
				m.input.SetValue(m.matchRes.Pattern)
				m.viewport.SetContent(m.matchRes.InputText)
			}
			m.updateRegexMatches()
		}
	}

	return m, cmd
}

func (m *model) updateRegexMatches() {
	if m.err = m.matchRes.FindMatches(); m.err != nil {
		return
	}
	m.matchRes.HighlightMatches(tsHighlight.Render)
	m.viewport.SetContent(m.matchRes.Highlighted)
}

func (m *model) getNextFocus() (tea.Model, tea.Cmd) {
	var nextFocus focus
	switch m.focus {
	case focusInput:
		nextFocus = focusContent
	case focusContent:
		nextFocus = focusCheatsheet
	case focusCheatsheet:
		nextFocus = focusInput
	default:
		nextFocus = focusInput
	}
	m.focus = nextFocus
	m.cheatsheet.SetDelegate(itemDelegate{focus: nextFocus})
	return m, nil
}

func (m *model) toggleState() (tea.Model, tea.Cmd) {
	switch m.state {
	case stateActive:
		m.state = stateExiting
	default:
		m.state = stateActive
	}
	return m, nil
}

func timeout(duration time.Duration) tea.Cmd {
	return func() tea.Msg {
		time.Sleep(duration)
		return tickMsg{}
	}
}
