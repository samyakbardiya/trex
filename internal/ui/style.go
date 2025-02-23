package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	colorsPerRow    = 8
	maxColors       = 16
	minWidth        = 80
	minHeight       = 24
	leftWidthRatio  = 0.70
	rightWidthRatio = 0.30
	minInputHeight  = 4
	minHelpHeight   = 2
	borderWidthDiff = 2  // self border
	widthDiff       = 20 // offset caused by the borders
	widthCheatsheet = 30 // width of the cheatsheet
)

var (
	cBlack        = lipgloss.Color("00")
	cRed          = lipgloss.Color("01")
	cGreen        = lipgloss.Color("02")
	cYellow       = lipgloss.Color("03")
	cBlue         = lipgloss.Color("04")
	cMagenta      = lipgloss.Color("05")
	cCyan         = lipgloss.Color("06")
	cLightGray    = lipgloss.Color("07")
	cGray         = lipgloss.Color("08")
	cLightRed     = lipgloss.Color("09")
	cLightGreen   = lipgloss.Color("10")
	cLightYellow  = lipgloss.Color("11")
	cLightBlue    = lipgloss.Color("12")
	cLightMagenta = lipgloss.Color("13")
	cLightCyan    = lipgloss.Color("14")
	cWhite        = lipgloss.Color("15")
)

// text-style
var (
	tsHelp                  = lipgloss.NewStyle().Foreground(cGray)
	tsHighlight             = lipgloss.NewStyle().Foreground(cBlack).Background(cGreen).Bold(true)
	tsNormal                = lipgloss.NewStyle()
	tsCheatsheetDescription = lipgloss.NewStyle()
	tsCheatsheetPattern     = lipgloss.NewStyle().Foreground(cYellow)
	tsListDefault           = lipgloss.NewStyle()
	tsListSelected          = lipgloss.NewStyle().Reverse(true).Bold(true)
)

// border-style
var (
	bsError   = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(cLightRed))
	bsFocus   = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(cLightBlue))
	bsUnfocus = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
	bsSuccess = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(cLightGreen))
)

func PreviewStyles() string {
	var b strings.Builder
	fmt.Fprintf(&b, "\n%s\n%s\n%s",
		bsError.Render("\tERROR\t"),
		bsFocus.Render("\tFOCUS\t"),
		bsUnfocus.Render("\tUNFOCUS\t"),
	)
	fmt.Fprintf(&b, "\n\n%s%s%s",
		tsHelp.Render("\tHELP\t"),
		tsHighlight.Render("\tHIGHLIGHT\t"),
		tsNormal.Render("\tNORMAL\t"),
	)
	return b.String()
}

func PreviewColors() string {
	var b strings.Builder
	b.WriteString("\n\n")
	renderColorPreview(&b, renderForegroundColor)
	b.WriteString("\n\n")
	renderColorPreview(&b, renderBackgroundColor)
	return b.String()
}

// renderColorPreview handles the common logic for rendering color previews
func renderColorPreview(b *strings.Builder, renderFunc func(int) string) {
	for i := 0; i < maxColors; i++ {
		b.WriteString(renderFunc(i))
		if (i+1)%colorsPerRow == 0 {
			b.WriteString("\n")
		}
	}
}

func renderForegroundColor(i int) string {
	return lipgloss.NewStyle().
		Foreground(lipgloss.Color(strconv.Itoa(i))).
		Render(fmt.Sprintf(" %3d ", i))
}

func renderBackgroundColor(i int) string {
	return lipgloss.NewStyle().
		Background(lipgloss.Color(strconv.Itoa(i))).
		Render(fmt.Sprintf(" %3d ", i))
}
