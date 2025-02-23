package ui

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/samyakbardiya/trex/internal/util"
)

// focus represents the current focus state of the UI.
type focus uint

const (
	focusInput focus = iota
	focusContent
)

// state represents the current state of the application
type state uint

const (
	stateActive state = iota
	stateAlertClipboard
	stateExiting
)

// tickMsg represents a tick event in the application
type tickMsg struct{}

type keyBinding struct {
	description string // description provides a human-readable explanation of the binding
	binding     string // binding represents the key sequence for this binding
}

type model struct {
	state    state            // current state of the application
	focus    focus            // current focus of the UI
	matchRes util.MatchResult // result of the regex matching operation
	input    textinput.Model  // model for handling input
	viewport viewport.Model   // model for handling content
	width    int              // width of the window
	height   int              // height of the window
	time     tickMsg          // tracks tick events for state transitions
	err      error            // any error encountered during application execution
}

func New(initialContent string) model {
	input := textinput.New()
	input.Placeholder = "RegEx"
	input.Focus()

	return model{
		state: stateActive,
		focus: focusInput,
		matchRes: util.MatchResult{
			InputText:   initialContent,
			Highlighted: initialContent,
		},
		input: input,
		err:   nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}
