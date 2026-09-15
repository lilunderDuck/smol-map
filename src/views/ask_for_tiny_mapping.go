package views

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const VIEW_ASK_FOR_TINY_MAPPING = 0

type AskForTinyMappingModel struct {
	// spinner  spinner.Model
	textInput textinput.Model
}

func NewAskForTinyMappingModel() AskForTinyMappingModel {
	ti := textinput.New()
	ti.Placeholder = "/path/to/tiny-v1-mapping.tiny"
	ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 156
	ti.SetWidth(20)

	return AskForTinyMappingModel{textInput: ti}
}

func (this AskForTinyMappingModel) Init() tea.Cmd {
	return textinput.Blink
}

func (this AskForTinyMappingModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "esc", "ctrl+c":
			return this, tea.Quit
		case "enter":
			return this, tea.Quit
		}
	}

	this.textInput, cmd = this.textInput.Update(msg)
	return this, cmd
}

var purple = lipgloss.Color("63")

func (this AskForTinyMappingModel) View() tea.View {
	var cursor *tea.Cursor
	if !this.textInput.VirtualCursor() {
		cursor = this.textInput.Cursor()
		cursor.Y += lipgloss.Height(this.headerView())
	}

	var sb strings.Builder
	fmt.Fprintf(&sb, "%s from %s\n")

	str := lipgloss.JoinVertical(lipgloss.Top, this.headerView(), this.textInput.View(), this.footerView())

	view := tea.NewView(str)
	view.Cursor = cursor
	return view
}

func (m AskForTinyMappingModel) headerView() string {
	return "What's your favorite Pokémon?\n"
}

func (m AskForTinyMappingModel) footerView() string {
	return "\n(esc to quit)"
}
