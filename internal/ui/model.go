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

type itemType uint

const (
	itemCheatcode itemType = iota
	itemTemplate
)

type item struct {
	itemType    itemType
	pattern     string
	description string
	testStr     string
}

func (i item) FilterValue() string { return "" }

var items = []list.Item{
	item{
		itemType:    itemTemplate,
		pattern:     "^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\\.[a-zA-Z]{2,}$",
		description: "Email Vali",
		testStr:     "test@example.com\nuser.name+tag@domain.co\njohn.doe@sub.example.org\nadmin@localhost\nuser@domain.com\ninvalid@domain\n@missingusername.com\nuser@.com\nuser@domain..com\nuser@domain.c",
	},
	item{
		itemType:    itemTemplate,
		pattern:     "^(https?|ftp):\\/\\/[^\\s/$.?#].[^\\s]*$",
		description: "URL Vali",
		testStr:     "http://example.com\nhttps://www.google.com\nftp://fileserver.com\nhttps://sub.domain.org/path?query=1\nhttp://localhost\ninvalid://url\nhttp:/missing.com\nhttps://\nftp:/invalid\nhttp://.com",
	},
	item{
		itemType:    itemTemplate,
		pattern:     "^\\+?[1-9]\\d{1,14}$",
		description: "Phone Vali",
		testStr:     "+1234567890\n1234567890\n+19876543210\n+447911123456\n987654321\n+1\n+123456789012345\n0123456789\n+1234567890123456\n+",
	},
	item{
		itemType:    itemTemplate,
		pattern:     "^(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)\\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)$",
		description: "IPv4 Vali",
		testStr:     "192.168.1.1\n255.255.255.255\n0.0.0.0\n127.0.0.1\n10.0.0.1\n256.256.256.256\n192.168.1\n192.168.1.256\n192.168.1.1.1\n192.168..1",
	},
	item{
		itemType:    itemTemplate,
		pattern:     "^([0-9a-fA-F]{1,4}:){7}[0-9a-fA-F]{1,4}$",
		description: "IPv6 Vali",
		testStr:     "2001:0db8:85a3:0000:0000:8a2e:0370:7334\n::1\nfe80::1ff:fe23:4567:890a\n2001:db8::\n2001:db8:0:0:0:0:2:1\n2001:db8:0:0:0:0:0:1\n2001:db8::1\n2001:db8:::1\n2001:db8:85a3::8a2e:370:7334\n::",
	},
	item{
		itemType:    itemTemplate,
		pattern:     "^\\d{4}-(0[1-9]|1[0-2])-(0[1-9]|[12]\\d|3[01])$",
		description: "Date Vali",
		testStr:     "2023-10-15\n1999-01-01\n2000-02-29\n2023-02-28\n2023-12-31\n2023-13-01\n2023-00-01\n2023-01-32\n2023-02-30\n2023-02-29",
	},
	item{
		itemType:    itemTemplate,
		pattern:     "^([01]\\d|2[0-3]):[0-5]\\d:[0-5]\\d$",
		description: "Time Vali",
		testStr:     "14:30:45\n00:00:00\n23:59:59\n12:34:56\n01:01:01\n24:00:00\n12:60:00\n12:00:60\n12:34\n123:45:67",
	},
	item{
		itemType:    itemTemplate,
		pattern:     "^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$",
		description: "MAC Vali",
		testStr:     "00:1A:2B:3C:4D:5E\n00-1A-2B-3C-4D-5E\n01:23:45:67:89:AB\nAA:BB:CC:DD:EE:FF\n00:00:00:00:00:00\n00:1A:2B:3C:4D\n00:1A:2B:3C:4D:5E:6F\n00:1A:2B:3C:4D:ZZ\n00:1A:2B:3C:4D:5\n00:1A:2B:3C:4D:5E:",
	},

	item{itemType: itemCheatcode, pattern: "\\n", description: "Newline"},
	item{itemType: itemCheatcode, pattern: "\\t", description: "Tab"},
	item{itemType: itemCheatcode, pattern: "\\r", description: "Carriage return"},
	item{itemType: itemCheatcode, pattern: ".", description: "Any single character"},
	item{itemType: itemCheatcode, pattern: "\\s", description: "Any whitespace character"},
	item{itemType: itemCheatcode, pattern: "\\S", description: "Any non-whitespace character"},
	item{itemType: itemCheatcode, pattern: "\\d", description: "Any digit"},
	item{itemType: itemCheatcode, pattern: "\\D", description: "Any non-digit"},
	item{itemType: itemCheatcode, pattern: "\\w", description: "Any word character"},
	item{itemType: itemCheatcode, pattern: "\\W", description: "Any non-word character"},
	item{itemType: itemCheatcode, pattern: "^", description: "Start of string"},
	item{itemType: itemCheatcode, pattern: "$", description: "End of string"},
	item{itemType: itemCheatcode, pattern: "\\b", description: "A word boundary"},
	item{itemType: itemCheatcode, pattern: "\\B", description: "Non-word boundary"},
	item{itemType: itemCheatcode, pattern: "a*", description: "Zero or more of a"},
	item{itemType: itemCheatcode, pattern: "a+", description: "One or more of a"},
	item{itemType: itemCheatcode, pattern: "a?", description: "Zero or one of a"},
	item{itemType: itemCheatcode, pattern: "a{3,6}", description: "Between 3 and 6 of a"},
	item{itemType: itemCheatcode, pattern: "[a-z]", description: "A character in the range: a-z"},
	item{itemType: itemCheatcode, pattern: "[a-zA-Z]", description: "A character in the range: a-z or A-Z"},
	item{itemType: itemCheatcode, pattern: "[[:digit:]]", description: "Decimal digits"},
	item{itemType: itemCheatcode, pattern: "[[:alnum:]]", description: "Letters and digits"},
	item{itemType: itemCheatcode, pattern: "[[:alpha:]]", description: "Letters"},
	item{itemType: itemCheatcode, pattern: "[[:lower:]]", description: "Lowercase letters"},
	item{itemType: itemCheatcode, pattern: "[[:upper:]]", description: "Uppercase letters"},
	item{itemType: itemCheatcode, pattern: "[[:space:]]", description: "Whitespace"},
	item{itemType: itemCheatcode, pattern: "[[:print:]]", description: "Visible characters"},
	item{itemType: itemCheatcode, pattern: "[[:graph:]]", description: "Visible characters (not space)"},
	item{itemType: itemCheatcode, pattern: "[[:punct:]]", description: "Visible punctuation characters"},
	item{itemType: itemCheatcode, pattern: "[[:xdigit:]]", description: "Hexadecimal digits"},
	item{itemType: itemCheatcode, pattern: "\\n", description: "Insert a newline"},
	item{itemType: itemCheatcode, pattern: "\\t", description: "Insert a tab"},
	item{itemType: itemCheatcode, pattern: "\\r", description: "Insert a carriage return"},
	item{itemType: itemCheatcode, pattern: "\\f", description: "Insert a form-feed"},
	item{itemType: itemCheatcode, pattern: "[abc]", description: "A single character of: a, b or c"},
	item{itemType: itemCheatcode, pattern: "[^abc]", description: "A character except: a, b or c"},
	item{itemType: itemCheatcode, pattern: "[^a-z]", description: "A character not in the range: a-z"},
	item{itemType: itemCheatcode, pattern: "a|b", description: "Alternate - match either a or b"},
	item{itemType: itemCheatcode, pattern: "(...)", description: "Capturing group"},
	item{itemType: itemCheatcode, pattern: "(?:...)", description: "Non-capturing group"},
	item{itemType: itemCheatcode, pattern: "(?P<name>...)", description: "Named Capturing Group"},
	item{itemType: itemCheatcode, pattern: "g", description: "Global"},
	item{itemType: itemCheatcode, pattern: "m", description: "Multiline"},
	item{itemType: itemCheatcode, pattern: "i", description: "Case insensitive"},
	item{itemType: itemCheatcode, pattern: "s", description: "Single line"},
	item{itemType: itemCheatcode, pattern: "a*", description: "Greedy quantifier"},
	item{itemType: itemCheatcode, pattern: "U", description: "Ungreedy"},
	item{itemType: itemCheatcode, pattern: "\\A", description: "Start of string"},
	item{itemType: itemCheatcode, pattern: "\\z", description: "Absolute end of string"},
	item{itemType: itemCheatcode, pattern: "\\xYY", description: "Hex character YY"},
	item{itemType: itemCheatcode, pattern: "\\x{YYYY}", description: "Hex character YYYY"},
	item{itemType: itemCheatcode, pattern: "\\ddd", description: "Octal character ddd"},
	item{itemType: itemCheatcode, pattern: "\\pX", description: "Unicode property X"},
	item{itemType: itemCheatcode, pattern: "\\p{...}", description: "Unicode property or script category"},
	item{itemType: itemCheatcode, pattern: "\\PX", description: "Negation of \\pX"},
	item{itemType: itemCheatcode, pattern: "\\P{...}", description: "Negation of \\p"},
	item{itemType: itemCheatcode, pattern: "\\Q...\\E", description: "Quote; treat as literals"},
	item{itemType: itemCheatcode, pattern: "[[:cntrl:]]", description: "Control characters"},
	item{itemType: itemCheatcode, pattern: "[[:ascii:]]", description: "ASCII codes 0-127"},
	item{itemType: itemCheatcode, pattern: "[[:blank:]]", description: "Space or tab only"},
	item{itemType: itemCheatcode, pattern: "\\v", description: "Vertical whitespace character"},
	item{itemType: itemCheatcode, pattern: "\\", description: "Makes any character literal"},
	item{itemType: itemCheatcode, pattern: "a{3}", description: "Exactly 3 of a"},
	item{itemType: itemCheatcode, pattern: "a{3,}", description: "3 or more of a"},
	item{itemType: itemCheatcode, pattern: "\\0", description: "Complete match contents"},
	item{itemType: itemCheatcode, pattern: "\\1", description: "Contents in capture group 1"},
	item{itemType: itemCheatcode, pattern: "$1", description: "Contents in capture group 1"},
	item{itemType: itemCheatcode, pattern: "${foo}", description: "Contents in capture group `foo"},
	item{itemType: itemCheatcode, pattern: "\\x{06fa}", description: "Hexadecimal replacement values"},
	item{itemType: itemCheatcode, pattern: "\\u06fa", description: "Hexadecimal replacement values"},
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

	var str []string
	switch i.itemType {
	case itemCheatcode:
		description := tsCheatsheetDescription.Render(i.description)
		pattern := tsCheatcode.Render(fmt.Sprintf(" %4s ", i.pattern))
		str = []string{pattern, description}
	case itemTemplate:
		pattern := tsTemplate.Render(fmt.Sprintf(" %4s ", i.pattern))
		description := tsCheatsheetDescription.Render(i.description)
		str = []string{description, pattern}
	}

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
