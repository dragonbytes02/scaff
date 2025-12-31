package app

import (
	"github.com/dragonbytes039/scaff/fst"
	"github.com/dragonbytes039/scaff/internal/msg"
	"github.com/dragonbytes039/scaff/internal/scripts"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *AppModel) runScript(selection scripts.ScriptApp, name_project string) tea.Cmd {

	return func() tea.Msg {

		template := scripts.TemplateName[selection]

		err := fst.SpreadTemplate(template, name_project)

		if err != nil {
			return msg.ScriptErrorMsg{Err: string("output") + ": " + err.Error()}
		}
		messagess := scripts.SuccessMsg(selection)
		return msg.ScriptFinishedMsg{Msg: messagess}
	}
}
