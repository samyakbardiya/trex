package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// text-style
var (
	tsHelp      = lipgloss.NewStyle().Foreground(lipgloss.Color("8")).Render
	tsHighlight = lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("2")).Bold(true).Render
	tsNormal    = lipgloss.NewStyle().Render
)

// border-style
var (
	bsError   = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("9")).Render
	bsFocus   = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("12")).Render
	bsUnfocus = lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).Render
)

func PreviewStyles() string {
	var b strings.Builder

	fmt.Fprint(&b,
		"\n"+bsError("\tERROR\t")+
			"\n"+bsFocus("\tFOCUS\t")+
			"\n"+bsUnfocus("\tUNFOCUS\t"),
	)
	fmt.Fprint(&b,
		"\n\n"+
			tsHelp("\tHELP\t")+
			tsHighlight("\tHIGHLIGHT\t")+
			tsNormal("\tNORMAL\t"),
	)
	return b.String()
}

func PreviewColors() string {
	const numColors = 128
	var b strings.Builder
	b.WriteString("\n\n")
	for i := 0; i < numColors; i++ {
		_if := fmt.Sprintf("%3d ", i)
		_is := strconv.Itoa(i)
		b.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color(_is)).Render(_if))
		if (i+1)%8 == 0 {
			b.WriteString("\n")
		}
	}
	b.WriteString("\n\n")
	for i := 0; i < numColors; i++ {
		_if := fmt.Sprintf(" %3d ", i)
		_is := strconv.Itoa(i)
		b.WriteString(lipgloss.NewStyle().Background(lipgloss.Color(_is)).Render(_if))
		if (i+1)%8 == 0 {
			b.WriteString("\n")
		}
	}
	return b.String()
}
