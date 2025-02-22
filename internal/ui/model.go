package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samyakbardiya/trex/internal/util"
)

type focus uint

const (
	focusInput focus = iota
	focusContent
)

type model struct {
	focus    focus
	matchRes util.MatchResult
	input    textinput.Model
	viewport viewport.Model
	err      error
}

func New(initialContent string) model {
	input := textinput.New()
	input.Placeholder = "RegEx"
	input.Focus()

	return model{
		input: input,
		focus: focusInput,
		matchRes: util.MatchResult{
			InputText:   initialContent,
			Highlighted: initialContent,
		},
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}
