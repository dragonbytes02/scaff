package Go

import (
	"dragonbytes02/scaff/internal/components/menu"
	"dragonbytes02/scaff/internal/msg"
	"dragonbytes02/scaff/internal/scripts"
)

var GoOptions = []menu.OptionsInfo{
	{Description: "Hexagonal Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.GoHexa}},
	{Description: "Modular Clean", Msg: msg.RunningScriptMsg{Script: scripts.GoMC}},
	{Description: "Model Controller Services", Msg: msg.RunningScriptMsg{Script: scripts.GoMcs}},
}

func InitialGoOptionModel() menu.Model {

	return menu.InitialModel(GoOptions, "select template")
}
