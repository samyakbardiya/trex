package ui

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type textStyle struct {
	help   lipgloss.Style
	match  lipgloss.Style
	normal lipgloss.Style
}

type borderStyle struct {
	error   lipgloss.Style
	focus   lipgloss.Style
	unfocus lipgloss.Style
}

func getBorderStyle() borderStyle {
	return borderStyle{
		error:   lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("9")),
		focus:   lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("12")),
		unfocus: lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()),
	}
}

func getTextStyle() textStyle {
	return textStyle{
		help:   lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
		match:  lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("2")).Bold(true),
		normal: lipgloss.NewStyle(),
	}
}

func PreviewStyles() string {
	var b strings.Builder
	bs := getBorderStyle()
	ts := getTextStyle()

	fmt.Fprint(&b,
		"\n"+bs.error.Render("\tERROR\t")+
			"\n"+bs.focus.Render("\tFOCUS\t")+
			"\n"+bs.unfocus.Render("\tUNFOCUS\t"),
	)
	fmt.Fprint(&b,
		"\n\n"+
			ts.help.Render("\tHELP\t")+
			ts.match.Render("\tMATCH\t")+
			ts.normal.Render("\tNORMAL\t"),
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
