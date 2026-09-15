package utils

import (
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetTerminalSize() (int, int) {
	// Verify if stdout is actually a terminal
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		panic("Not running in a terminal")
	}

	// Fetch width and height
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		panic(fmt.Errorf("Error getting terminal size: %v\n", err))
	}

	return width - 10, height
}
