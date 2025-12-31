package msg

import "github.com/dragonbytes039/scaff/internal/scripts"

type ChangeToLanguageMsg struct{}

//js
type ChangeToParadigmMsg struct{}
type ChangeToJsPooMsg struct{}
type ChangeToJsFuntionalMsg struct{}

//go
type ChangeToGoOptionsMsg struct{}

//scripts
type RunningScriptMsg struct {
	Script scripts.ScriptApp
}

type ScriptFinishedMsg struct {
	Msg string
}
type ScriptErrorMsg struct {
	Err string
}
