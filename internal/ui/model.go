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
	item{description: "Newline", pattern: "\\n"},
	item{description: "Carriage return", pattern: "\\r"},
	item{description: "Tab", pattern: "\\t"},
	item{description: "A single character of: a, b or c", pattern: "[abc]"},
	item{description: "A character except: a, b or c", pattern: "[^abc]"},
	item{description: "A character in the range: a-z", pattern: "[a-z]"},
	item{description: "A character not in the range: a-z", pattern: "[^a-z]"},
	item{description: "A character in the range: a-z or A-Z", pattern: "[a-zA-Z]"},
	item{description: "Letters and digits", pattern: "[[:alnum:]]"},
	item{description: "Letters", pattern: "[[:alpha:]]"},
	item{description: "ASCII codes 0-127", pattern: "[[:ascii:]]"},
	item{description: "Space or tab only", pattern: "[[:blank:]]"},
	item{description: "Control characters", pattern: "[[:cntrl:]]"},
	item{description: "Decimal digits", pattern: "[[:digit:]]"},
	item{description: "Visible characters (not space)", pattern: "[[:graph:]]"},
	item{description: "Lowercase letters", pattern: "[[:lower:]]"},
	item{description: "Visible characters", pattern: "[[:print:]]"},
	item{description: "Visible punctuation characters", pattern: "[[:punct:]]"},
	item{description: "Whitespace", pattern: "[[:space:]]"},
	item{description: "Uppercase letters", pattern: "[[:upper:]]"},
	item{description: "Word characters", pattern: "[[:word:]]"},
	item{description: "Hexadecimal digits", pattern: "[[:xdigit:]]"},
	item{description: "Any single character", pattern: "."},
	item{description: "Alternate - match either a or b", pattern: "a|b"},
	item{description: "Any whitespace character", pattern: "\\s"},
	item{description: "Any non-whitespace character", pattern: "\\S"},
	item{description: "Any digit", pattern: "\\d"},
	item{description: "Any non-digit", pattern: "\\D"},
	item{description: "Any word character", pattern: "\\w"},
	item{description: "Any non-word character", pattern: "\\W"},
	item{description: "Vertical whitespace character", pattern: "\\v"},
	item{description: "Unicode property X", pattern: "\\pX"},
	item{description: "Unicode property or script category", pattern: "\\p{...}"},
	item{description: "Negation of \\pX", pattern: "\\PX"},
	item{description: "Negation of \\p", pattern: "\\P{...}"},
	item{description: "Quote; treat as literals", pattern: "\\Q...\\E"},
	item{description: "Hex character YY", pattern: "\\xYY"},
	item{description: "Hex character YYYY", pattern: "\\x{YYYY}"},
	item{description: "Octal character ddd", pattern: "\\ddd"},
	item{description: "Makes any character literal", pattern: "\\"},
	item{description: "Non-capturing group", pattern: "(?:...)"},
	item{description: "Capturing group", pattern: "(...)"},
	item{description: "Named Capturing Group", pattern: "(?P<name>...)"},
	item{description: "Zero or one of a", pattern: "a?"},
	item{description: "Zero or more of a", pattern: "a*"},
	item{description: "One or more of a", pattern: "a+"},
	item{description: "Exactly 3 of a", pattern: "a{3}"},
	item{description: "3 or more of a", pattern: "a{3,}"},
	item{description: "Between 3 and 6 of a", pattern: "a{3,6}"},
	item{description: "Greedy quantifier", pattern: "a*"},
	item{description: "Start of string", pattern: "^"},
	item{description: "End of string", pattern: "$"},
	item{description: "Start of string", pattern: "\\A"},
	item{description: "Absolute end of string", pattern: "\\z"},
	item{description: "A word boundary", pattern: "\\b"},
	item{description: "Non-word boundary", pattern: "\\B"},
	item{description: "Global", pattern: "g"},
	item{description: "Multiline", pattern: "m"},
	item{description: "Case insensitive", pattern: "i"},
	item{description: "Single line", pattern: "s"},
	item{description: "Ungreedy", pattern: "U"},
	item{description: "Complete match contents", pattern: "\\0"},
	item{description: "Contents in capture group 1", pattern: "\\1"},
	item{description: "Contents in capture group 1", pattern: "$1"},
	item{description: "Contents in capture group `foo", pattern: "${foo}"},
	item{description: "Hexadecimal replacement values", pattern: "\\x{06fa}"},
	item{description: "Hexadecimal replacement values", pattern: "\u06fa"},
	item{description: "Insert a tab", pattern: "\\t"},
	item{description: "Insert a carriage return", pattern: "\\r"},
	item{description: "Insert a newline", pattern: "\\n"},
	item{description: "Insert a form-feed", pattern: "\\f"},
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
	str := []string{description, pattern}

	width := m.Width()
	style := tsListDefault
	if d.focus == focusCheatsheet && index == m.Index() {
		style = tsListSelected
	}
	fmt.Fprint(w, style.Width(width).MaxWidth(width).Render(str...))
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
