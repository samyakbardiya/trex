package ui

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

type TextStyle struct {
	help   lipgloss.Style
	match  lipgloss.Style
	normal lipgloss.Style
}

type BorderStyle struct {
	error   lipgloss.Style
	focus   lipgloss.Style
	unfocus lipgloss.Style
}

func GetBorderStyle() BorderStyle {
	return BorderStyle{
		error:   lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("9")),
		focus:   lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()).BorderForeground(lipgloss.Color("12")),
		unfocus: lipgloss.NewStyle().BorderStyle(lipgloss.RoundedBorder()),
	}
}

func GetTextStyle() TextStyle {
	return TextStyle{
		help:   lipgloss.NewStyle().Foreground(lipgloss.Color("8")),
		match:  lipgloss.NewStyle().Foreground(lipgloss.Color("0")).Background(lipgloss.Color("2")).Bold(true),
		normal: lipgloss.NewStyle(),
	}
}

func TestAllStyle() string {
	s := ""
	bs := GetBorderStyle()
	ts := GetTextStyle()

	s += fmt.Sprint(
		"\n" + bs.error.Render("\tERROR\t") +
			"\n" + bs.focus.Render("\tFOCUS\t") +
			"\n" + bs.unfocus.Render("\tUNFOCUS\t"),
	)

	s += "\n\n"
	s += fmt.Sprint(
		ts.help.Render("\tHELP\t") +
			ts.match.Render("\tMATCH\t") +
			ts.normal.Render("\tNORMAL\t"),
	)

	return s
}

func TestAllColors() string {
	const colorBit = 128
	s := ""

	s += "\n\n"
	for i := 0; i < colorBit; i++ {
		_if := fmt.Sprintf("%3d ", i)
		_is := strconv.Itoa(i)
		s += lipgloss.NewStyle().Foreground(lipgloss.Color(_is)).Render(_if)
		if (i+1)%8 == 0 {
			s += "\n"
		}
	}

	s += "\n\n"
	for i := 0; i < colorBit; i++ {
		_if := fmt.Sprintf(" %3d ", i)
		_is := strconv.Itoa(i)
		s += lipgloss.NewStyle().Background(lipgloss.Color(_is)).Render(_if)
		if (i+1)%8 == 0 {
			s += "\n"
		}
	}

	return s
}
