package language

import (
	"github.com/dragonbytes039/scaff/internal/components/menu"
	"github.com/dragonbytes039/scaff/internal/msg"

	"github.com/charmbracelet/lipgloss"
)

var LanguagesOptions = []menu.OptionsInfo{
	{Description: "Go", Msg: msg.ChangeToGoOptionsMsg{}, Color: lipgloss.NewStyle().Foreground(lipgloss.Color("#00ADD8"))},
	{Description: "Typescript", Msg: msg.ChangeToParadigmMsg{}, Color: lipgloss.NewStyle().Foreground(lipgloss.Color("#0B64D9"))},
}

func InitialLanguageModel() menu.Model {

	return menu.InitialModel(LanguagesOptions, "select language")
}
