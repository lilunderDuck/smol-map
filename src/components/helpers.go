package components

import (
	"fmt"

	"charm.land/lipgloss/v2"
)

const MAX_PAGE_WIDTH = 101

var RoundedBorderBox = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderTop(true).
	BorderLeft(true).
	BorderBottom(true).
	BorderRight(true).
	BorderForeground(lipgloss.Color("243")).
	Width(MAX_PAGE_WIDTH)

var CenterAligned = lipgloss.NewStyle().
	AlignHorizontal(lipgloss.Center).
	Width(MAX_PAGE_WIDTH)

// ...

func Renderf(style lipgloss.Style, stuff string, format ...any) string {
	return style.Render(fmt.Sprintf(stuff, format...))
}

func Render(style lipgloss.Style, stuff string) string {
	return style.Render(stuff)
}
