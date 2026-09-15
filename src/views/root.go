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
			VIEW_SEARCH_NAME:          NewSearchNameModel(),
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
	return this.viewMap[this.currentView].Update(msg)
}

func (this RootModel) View() tea.View {
	return this.viewMap[this.currentView].View()
}
