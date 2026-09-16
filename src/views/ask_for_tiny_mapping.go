package views

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"smolmap/src/components"

	// "smolmap/src/tiny"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const VIEW_ASK_FOR_TINY_MAPPING = 0

type AskForTinyMappingModel struct {
	// spinner  spinner.Model
	textInput      textinput.Model
	textInputError error
}

func NewAskForTinyMappingModel() AskForTinyMappingModel {
	ti := textinput.New()
	ti.Placeholder = "Example: /path/to/tiny-v1-mapping.tiny"
	// ti.SetVirtualCursor(false)
	ti.Focus()
	ti.CharLimit = 100
	ti.SetWidth(100)
	ti.Validate = func(inputPath string) error {
		if inputPath == "" {
			return fmt.Errorf("You've provided an empty path...")
		}

		// return tiny.DetectForTinyV1(inputPath, nil)
		return nil
	}

	return AskForTinyMappingModel{
		textInput: ti,
	}
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
			validationErr := this.textInput.Validate(this.textInput.Value())
			if validationErr != nil {
				this.textInputError = validationErr
				break
			}
			return this, tea.Quit
		}
	}

	this.textInput, cmd = this.textInput.Update(msg)
	return this, cmd
}

func (this AskForTinyMappingModel) View() tea.View {
	var cursor *tea.Cursor
	if !this.textInput.VirtualCursor() {
		cursor = this.textInput.Cursor()
		cursor.Y += lipgloss.Height(this.headerView())
	}

	everything := components.RoundedBorderBox.Render(
		lipgloss.JoinVertical(
			lipgloss.Top,
			this.headerView(),
			this.textInput.View(),
			this.footerView(),
		),
	)

	view := tea.NewView(everything)
	view.Cursor = cursor
	return view
}

func (this AskForTinyMappingModel) headerView() string {
	var sb strings.Builder
	fmt.Fprint(&sb, components.Render(
		components.CenterAligned,
		components.LOGO,
	))
	fmt.Fprintf(&sb, "\nEnter your .tiny v1 mapping path:")

	return sb.String()
}

func (this AskForTinyMappingModel) footerView() string {
	var sb strings.Builder
	if this.textInputError != nil {
		fmt.Fprintf(&sb, "%s%v%s\n\n", components.COLOR_RED, this.textInputError, components.F_RESET)
	} else {
		fmt.Fprintf(&sb, "\n\n")
	}

	components.RenderShortcutHint(&sb, this.footerViewShortcutMap())

	return sb.String()
}

func (this AskForTinyMappingModel) footerViewShortcutMap() []components.ShortcutHint {
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
