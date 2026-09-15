package components

import (
	"fmt"
	"image/color"

	"charm.land/lipgloss/v2"
)

var anotherStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	BorderTop(true).
	BorderLeft(true).
	BorderBottom(true).
	BorderRight(true).
	PaddingLeft(2).
	PaddingRight(2)

func Boxf(borderColor color.Color, stuff string, format ...any) string {
	return anotherStyle.BorderForeground(borderColor).Render(fmt.Sprintf(stuff, format...))
}
