package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
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
	focusCheatsheet
)

// state represents the current state of the application
type state uint

const (
	stateActive state = iota
	stateNotification
	stateExiting
)

// tickMsg represents a tick event in the application
type tickMsg struct{}

type keyBinding struct {
	description string // description provides a human-readable explanation of the binding
	binding     string // binding represents the key sequence for this binding
}

type item struct {
	description string
	pattern     string
}

func (i item) FilterValue() string { return "" }

var items = []list.Item{
	item{pattern: "\\n", description: "Newline"},
	item{pattern: "\\t", description: "Tab"},
	item{pattern: "\\r", description: "Carriage return"},
	item{pattern: ".", description: "Any single character"},
	item{pattern: "\\s", description: "Any whitespace character"},
	item{pattern: "\\S", description: "Any non-whitespace character"},
	item{pattern: "\\d", description: "Any digit"},
	item{pattern: "\\D", description: "Any non-digit"},
	item{pattern: "\\w", description: "Any word character"},
	item{pattern: "\\W", description: "Any non-word character"},
	item{pattern: "^", description: "Start of string"},
	item{pattern: "$", description: "End of string"},
	item{pattern: "\\b", description: "A word boundary"},
	item{pattern: "\\B", description: "Non-word boundary"},
	item{pattern: "a*", description: "Zero or more of a"},
	item{pattern: "a+", description: "One or more of a"},
	item{pattern: "a?", description: "Zero or one of a"},
	item{pattern: "a{3,6}", description: "Between 3 and 6 of a"},
	item{pattern: "[a-z]", description: "A character in the range: a-z"},
	item{pattern: "[a-zA-Z]", description: "A character in the range: a-z or A-Z"},
	item{pattern: "[[:digit:]]", description: "Decimal digits"},
	item{pattern: "[[:alnum:]]", description: "Letters and digits"},
	item{pattern: "[[:alpha:]]", description: "Letters"},
	item{pattern: "[[:lower:]]", description: "Lowercase letters"},
	item{pattern: "[[:upper:]]", description: "Uppercase letters"},
	item{pattern: "[[:space:]]", description: "Whitespace"},
	item{pattern: "[[:print:]]", description: "Visible characters"},
	item{pattern: "[[:graph:]]", description: "Visible characters (not space)"},
	item{pattern: "[[:punct:]]", description: "Visible punctuation characters"},
	item{pattern: "[[:xdigit:]]", description: "Hexadecimal digits"},
	item{pattern: "\\n", description: "Insert a newline"},
	item{pattern: "\\t", description: "Insert a tab"},
	item{pattern: "\\r", description: "Insert a carriage return"},
	item{pattern: "\\f", description: "Insert a form-feed"},
	item{pattern: "[abc]", description: "A single character of: a, b or c"},
	item{pattern: "[^abc]", description: "A character except: a, b or c"},
	item{pattern: "[^a-z]", description: "A character not in the range: a-z"},
	item{pattern: "a|b", description: "Alternate - match either a or b"},
	item{pattern: "(...)", description: "Capturing group"},
	item{pattern: "(?:...)", description: "Non-capturing group"},
	item{pattern: "(?P<name>...)", description: "Named Capturing Group"},
	item{pattern: "g", description: "Global"},
	item{pattern: "m", description: "Multiline"},
	item{pattern: "i", description: "Case insensitive"},
	item{pattern: "s", description: "Single line"},
	item{pattern: "a*", description: "Greedy quantifier"},
	item{pattern: "U", description: "Ungreedy"},
	item{pattern: "\\A", description: "Start of string"},
	item{pattern: "\\z", description: "Absolute end of string"},
	item{pattern: "\\xYY", description: "Hex character YY"},
	item{pattern: "\\x{YYYY}", description: "Hex character YYYY"},
	item{pattern: "\\ddd", description: "Octal character ddd"},
	item{pattern: "\\pX", description: "Unicode property X"},
	item{pattern: "\\p{...}", description: "Unicode property or script category"},
	item{pattern: "\\PX", description: "Negation of \\pX"},
	item{pattern: "\\P{...}", description: "Negation of \\p"},
	item{pattern: "\\Q...\\E", description: "Quote; treat as literals"},
	item{pattern: "[[:cntrl:]]", description: "Control characters"},
	item{pattern: "[[:ascii:]]", description: "ASCII codes 0-127"},
	item{pattern: "[[:blank:]]", description: "Space or tab only"},
	item{pattern: "\\v", description: "Vertical whitespace character"},
	item{pattern: "\\", description: "Makes any character literal"},
	item{pattern: "a{3}", description: "Exactly 3 of a"},
	item{pattern: "a{3,}", description: "3 or more of a"},
	item{pattern: "\\0", description: "Complete match contents"},
	item{pattern: "\\1", description: "Contents in capture group 1"},
	item{pattern: "$1", description: "Contents in capture group 1"},
	item{pattern: "${foo}", description: "Contents in capture group `foo"},
	item{pattern: "\\x{06fa}", description: "Hexadecimal replacement values"},
	item{pattern: "\\u06fa", description: "Hexadecimal replacement values"},
}

type itemDelegate struct {
	focus focus
}

func (d itemDelegate) Height() int                             { return 1 }
func (d itemDelegate) Spacing() int                            { return 0 }
func (d itemDelegate) Update(_ tea.Msg, _ *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	description := tsCheatsheetDescription.Render(i.description)
	pattern := tsCheatsheetPattern.Render(fmt.Sprintf(" %4s ", i.pattern))
	str := []string{pattern, description}

	width := m.Width()
	style := tsListDefault
	if d.focus == focusCheatsheet && index == m.Index() {
		style = tsListSelected
	}
	fmt.Fprint(w, style.MaxWidth(width).Render(str...))
}

type model struct {
	state      state            // current state of the application
	focus      focus            // current focus of the UI
	matchRes   util.MatchResult // result of the regex matching operation
	input      textinput.Model  // model for handling input
	viewport   viewport.Model   // model for handling content
	cheatsheet list.Model       // model for handling cheatsheet
	width      int              // width of the window
	height     int              // height of the window
	time       tickMsg          // tracks tick events for state transitions
	err        error            // any error encountered during application execution
}

func New(initialContent string) model {
	in := textinput.New()
	in.Placeholder = "RegEx"
	in.Prompt = ""
	in.Focus()

	ch := list.New(items, itemDelegate{focus: focusInput}, minWidth*rightWidthRatio, 48)
	ch.SetFilteringEnabled(false)
	ch.SetShowFilter(false)
	ch.SetShowHelp(false)
	ch.SetShowStatusBar(false)
	ch.SetShowTitle(false)
	ch.Styles.PaginationStyle = tsNormal
	ch.Styles.HelpStyle = tsHelp

	vp := viewport.New(minWidth*leftWidthRatio, minHeight)
	vp.SetContent(initialContent)

	return model{
		state: stateActive,
		focus: focusInput,
		matchRes: util.MatchResult{
			InputText:   initialContent,
			Highlighted: initialContent,
		},
		input:      in,
		viewport:   vp,
		cheatsheet: ch,
		err:        nil,
	}
}

func (m model) Init() tea.Cmd {
	return textinput.Blink
}
