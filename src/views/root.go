package views

// A simple program demonstrating the spinner component from the Bubbles
// component library.

import (
	tea "charm.land/bubbletea/v2"
)

type RootModel struct {
	currentView int
	viewMap     map[int]tea.Model
}

func NewRootModel() RootModel {
	return RootModel{
		viewMap: map[int]tea.Model{
			VIEW_ASK_FOR_TINY_MAPPING: NewAskForTinyMappingModel(),
		},
		currentView: VIEW_ASK_FOR_TINY_MAPPING,
	}
}

func (this *RootModel) GoTo(view int) {
	this.currentView = view
}

func (this RootModel) Init() tea.Cmd {
	return this.viewMap[this.currentView].Init()
}

func (this RootModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			// this.quitting = true
			return this, tea.Quit
		default:
			return this, nil
		}

	default:
		return this.viewMap[this.currentView].Update(msg)
	}
}

func (this RootModel) View() tea.View {
	return this.viewMap[this.currentView].View()
}
