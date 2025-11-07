package msg

import "dragonbytes02/scaff/internal/scripts"

type ChangeToLanguageMsg struct{}
type ChangeToParadigmMsg struct{}
type ChangeToJsPooMsg struct{}
type ChangeToJsFuntionalMsg struct{}

type RunningScriptMsg struct {
	Script scripts.ScriptApp
}

type ScriptFinishedMsg struct {
	Msg string
}
type ScriptErrorMsg struct {
	Err string
}
