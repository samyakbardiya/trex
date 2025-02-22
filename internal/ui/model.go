package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	CurrFocus uint
	ErrMsg    error
)

const (
	TextInput CurrFocus = iota
	ContentView
)

var (
	modelStyle = lipgloss.NewStyle().
			Width(80).
			Height(5).
			Align(lipgloss.Center, lipgloss.Center).
			BorderStyle(lipgloss.HiddenBorder())
	focusedModelStyle = lipgloss.NewStyle().
				Width(80).
				Height(5).
				Align(lipgloss.Center, lipgloss.Center).
				BorderStyle(lipgloss.NormalBorder()).
				BorderForeground(lipgloss.Color("69"))
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type Model struct {
	state       CurrFocus
	err         error
	content     string
	highlighted string
	expr        string
	matches     [][]int
	textInput   textinput.Model
	viewport    viewport.Model
}

func InitialModel(content string) Model {
	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.CharLimit = 156
	ti.Width = 20
	ti.Focus()

	return Model{
		textInput: ti,
		state:     TextInput,
		err:       nil,
		content:   string(content),
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m Model) currentFocusedModel() string {
	if m.state == TextInput {
		return "textInput"
	}
	return "content"
}
