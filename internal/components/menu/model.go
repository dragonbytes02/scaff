package menu

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type OptionsInfo struct {
	Description string
	Msg         tea.Msg
	Color       lipgloss.Style
}

type Model struct {
	choices []OptionsInfo
	cursor  int
	title   string
}

func InitialModel(options []OptionsInfo, title string) Model {

	return Model{
		choices: options,
		cursor:  0,
		title:   title,
	}
}
