package components

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

var keyShortcutColor = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffa65e"))
var keyShortcutSeperatorColor = lipgloss.NewStyle().Foreground(lipgloss.BrightBlack)

type ShortcutHint struct {
	Keys        []string
	Description string
}

func RenderShortcutHint(sb *strings.Builder, shortcutMap []ShortcutHint) {
	totalLen := len(shortcutMap)
	for shortcutIndex, currentShortcut := range shortcutMap {
		keyLen := len(currentShortcut.Keys)
		for keyIndex, key := range currentShortcut.Keys {
			fmt.Fprint(sb, keyShortcutColor.Render(key))

			includeExtraSlashes := keyIndex != keyLen-1
			if includeExtraSlashes {
				fmt.Fprint(sb, keyShortcutSeperatorColor.Render("/"))
			}
		}

		fmt.Fprint(sb, " ", currentShortcut.Description)

		includeExtraDot := shortcutIndex != totalLen-1
		if includeExtraDot {
			fmt.Fprint(sb, keyShortcutSeperatorColor.Render(" • "))
		}
	}
}
