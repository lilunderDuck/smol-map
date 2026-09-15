package components

import (
	"fmt"
	"strings"
)

func MappingName(mappingName string, className string, classPath string) string {
	// _, _ := utils.GetTerminalSize()

	var sb strings.Builder
	fmt.Fprintf(&sb, "%s from %s\n",
		format(className, COLOR_MAGENTA, STYLE_BOLD),
		format(classPath, COLOR_BLUE, STYLE_BOLD),
	)

	fmt.Fprintf(&sb, "╰ %s : %s\n",
		format("Official (obfuscated) name", COLOR_GRAY),
		format("aa", COLOR_CYAN, STYLE_BOLD),
	)

	fmt.Fprintf(&sb, "╰ %s : %s",
		format("Intermediary              ", COLOR_GRAY),
		format("net/minecraft/class_155", COLOR_CYAN, STYLE_BOLD),
	)

	return sb.String()
}
