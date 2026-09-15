package components

import (
	"fmt"
	"strings"
)

const (
	F_RESET       = "\x1b[0m"
	COLOR_RED     = "\x1b[31m"
	COLOR_GREEN   = "\x1b[32m"
	COLOR_YELLOW  = "\x1b[33m"
	COLOR_BLUE    = "\x1b[34m"
	COLOR_MAGENTA = "\x1b[35m"
	COLOR_CYAN    = "\x1b[36m"
	COLOR_WHITE   = "\x1b[37m"
	COLOR_GRAY    = "\033[90m"
)

const (
	STYLE_BOLD          = "\x1B[1m"
	STYLE_ITALIC        = "\x1B[3m"
	STYLE_UNDERLINE     = "\x1B[4m"
	STYLE_STRIKETHROUGH = "\x1B[9m"
	STYLE_NONE          = ""
)

func format(something string, format ...string) string {
	return fmt.Sprintf("%s%s%s", strings.Join(format, ""), something, F_RESET)
}

func repeat(something string, count int) string {
	return strings.Repeat(something, count)
}
