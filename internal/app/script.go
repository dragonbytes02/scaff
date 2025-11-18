package app

import (
	"dragonbytes02/scaff/internal/msg"
	"dragonbytes02/scaff/internal/scripts"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

func (m *AppModel) runScript(selection scripts.ScriptApp, name_project string) tea.Cmd {

	return func() tea.Msg {

		time.Sleep(3 * time.Second)

		url := scripts.ScriptsDetailsUrls[selection]

		script := scripts.InitialScript(name_project, url)

		err := script.Execute()

		if err != nil {
			return msg.ScriptErrorMsg{Err: string("output") + ": " + err.Error()}
		}
		messagess := scripts.SuccessMsg(selection)
		return msg.ScriptFinishedMsg{Msg: messagess}
	}
}
