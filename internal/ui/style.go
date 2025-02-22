package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	cBlack       = "0"
	cGreen       = "2"
	cGray        = "8"
	cRed         = "9"
	cBlue        = "12"
	colorsPerRow = 8
	maxColors    = 16
)

// text-style
var (
	tsHelp      = lipgloss.NewStyle().Foreground(lipgloss.Color(cGray))
	tsHighlight = lipgloss.NewStyle().Foreground(lipgloss.Color(cBlack)).Background(lipgloss.Color(cGreen)).Bold(true)
	tsNormal    = lipgloss.NewStyle()
)

// border-style
var (
	bsError   = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(cRed))
	bsFocus   = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color(cBlue))
	bsUnfocus = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder())
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
