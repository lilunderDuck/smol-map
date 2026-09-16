package views

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	"fmt"
	"smolmap/src/components"
	"smolmap/src/tiny"
	"strings"

	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

const VIEW_SEARCH_NAME = 1

type SearchNameModel struct {
	textInput     textinput.Model
	mapping       *tiny.MappingDatabase
	mappingResult []tiny.Mapping
	isHintShown   bool
}

var EMPTY_RESULT = []tiny.Mapping{}

func NewSearchNameModel(tinyMappingFile string) SearchNameModel {
	mapping, err := tiny.ParseTiny(tinyMappingFile)
	if err != nil {
		panic(err)
	}

	ti := textinput.New()
	ti.Placeholder = "Search for any class, method or field."
	ti.Focus()
	ti.CharLimit = 100
	ti.SetWidth(components.MAX_PAGE_WIDTH)

	return SearchNameModel{
		textInput:     ti,
		mapping:       mapping,
		mappingResult: EMPTY_RESULT,
		isHintShown:   false,
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
			this.searchMapping()
		case "/":
			this.isHintShown = !this.isHintShown
			return this, cmd
		}
	}

	this.textInput, cmd = this.textInput.Update(msg)
	return this, cmd
}

func (this *SearchNameModel) searchMapping() {
	mappingResult, ok := this.mapping.Lookup[this.textInput.Value()]
	if ok {
		this.mappingResult = mappingResult
	} else {
		this.mappingResult = EMPTY_RESULT
	}
}

func (this SearchNameModel) View() tea.View {
	var cursor *tea.Cursor
	if !this.textInput.VirtualCursor() {
		cursor = this.textInput.Cursor()
		cursor.Y += lipgloss.Height(this.headerView())
	}

	everything := components.RoundedBorderBox.
		PaddingLeft(1).
		PaddingRight(1).
		Render(lipgloss.JoinVertical(
			lipgloss.Top,
			this.headerView(),
			this.searchView(),
			this.footerView(),
		))

	view := tea.NewView(everything)
	view.Cursor = cursor
	return view
}

func (this SearchNameModel) headerView() string {
	return this.textInput.View()
}

var searchContentView = lipgloss.NewStyle().
	Width(components.MAX_PAGE_WIDTH).
	Height(15)

func (this SearchNameModel) searchView() string {
	var sb strings.Builder

	fmt.Fprintln(&sb)
	if this.isHintShown {
		fmt.Fprintln(&sb, "If you want to search nested class, you can search like this: class_2841$class_6563\n")
	} else {
		if len(this.mappingResult) != 0 {
			components.RenderMappingName(&sb, &this.mappingResult[0], this.mapping)
		} else {
			fmt.Fprintf(&sb, "Nothing here...\n")
		}
		fmt.Fprintln(&sb)
	}

	return searchContentView.Render(sb.String())
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
			Description: "confirm search",
		},
		{
			Keys:        []string{"/"},
			Description: "toggle hint",
		},
	}
}
