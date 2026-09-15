package views

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"smolmap/src/components"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const VIEW_SEARCH_NAME = 1

type SearchNameModel struct {
	// spinner  spinner.Model
	textInput textinput.Model
}

func NewSearchNameModel() SearchNameModel {
	ti := textinput.New()
	ti.Placeholder = "Search for something, example: class_1031"
	ti.Focus()
	ti.CharLimit = 100
	ti.SetWidth(100)

	return SearchNameModel{
		textInput: ti,
	}
}

func (this SearchNameModel) Init() tea.Cmd {
	return textinput.Blink
}

func (this SearchNameModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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

func (this SearchNameModel) View() tea.View {
	var cursor *tea.Cursor
	if !this.textInput.VirtualCursor() {
		cursor = this.textInput.Cursor()
		cursor.Y += lipgloss.Height(this.headerView())
	}

	everything := components.Render(
		components.RoundedBorderBox.
			PaddingLeft(1).
			PaddingRight(1).
			BorderForeground(purple),
		lipgloss.JoinVertical(lipgloss.Top, this.headerView(), this.searchView(), this.footerView()),
	)

	view := tea.NewView(everything)
	view.Cursor = cursor
	return view
}

func (this SearchNameModel) headerView() string {
	// var sb strings.Builder
	// return sb.String()
	return this.textInput.View()
}

func (this SearchNameModel) searchView() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "")
	fmt.Fprintln(&sb, components.MappingName("class_155", "SharedConstants", "net/minecraft/SharedConstants"))
	fmt.Fprintln(&sb, "")
	return sb.String()
}

func (this SearchNameModel) footerView() string {
	var sb strings.Builder
	fmt.Fprintln(&sb, "")
	components.RenderShortcutHint(&sb, this.footerViewShortcutMap())
	return sb.String()
}

func (this SearchNameModel) footerViewShortcutMap() []components.ShortcutHint {
	return []components.ShortcutHint{
		{
			Keys:        []string{"ESC", "Ctrl+C"},
			Description: "quit",
		},
		{
			Keys:        []string{"Enter"},
			Description: "confirm",
		},
	}
}
