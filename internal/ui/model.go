package ui

import (
	"time"

	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type (
	SessionState uint
	ErrMsg       error
)

const (
	DefaultTime              = time.Minute
	TextInput   SessionState = iota
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

type MainModel struct {
	state     SessionState
	textInput textinput.Model
	err       error
	viewport  viewport.Model
	content   string
}

func InitialModel(timeout time.Duration, content string) MainModel {

	ti := textinput.New()
	ti.Placeholder = "Pikachu"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return MainModel{
		textInput: ti,
		err:       nil,
		content:   string(content),
	}
}

func (m MainModel) Init() tea.Cmd {
	return tea.Batch(textinput.Blink)
}

func (m MainModel) currentFocusedModel() string {
	if m.state == TextInput {
		return "textInput"
	}
	return "content"
}
