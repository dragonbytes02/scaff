package Go

import (
	"dragonbytes02/scaff/internal/components/menu"
	"dragonbytes02/scaff/internal/msg"
	"dragonbytes02/scaff/internal/scripts"
)

var GoOptions = []menu.OptionsInfo{
	{Description: "Clean Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.GoClean}},
	{Description: "Hexagonal Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.GoHexa}},
	{Description: "Modular Clean", Msg: msg.RunningScriptMsg{Script: scripts.GoMC}},
	{Description: "ELM", Msg: msg.RunningScriptMsg{Script: scripts.GoEML}},
}

func InitialGoOptionModel() menu.Model {

	return menu.InitialModel(GoOptions, "select template")
}
