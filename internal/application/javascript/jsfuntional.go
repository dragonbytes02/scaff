package javascript

import (
	"dragonbytes02/scaff/internal/components/menu"
	"dragonbytes02/scaff/internal/msg"
	"dragonbytes02/scaff/internal/scripts"
)

var jsfuntionalOptions = []menu.OptionsInfo{
	{Description: "Clean Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.JsFuntionalClean}},
	{Description: "Hexagonal Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.JsFuntionalHexa}},
	{Description: "Model Controller Services", Msg: msg.RunningScriptMsg{Script: scripts.JsFuntionalMcs}},
	{Description: "Modular Clean", Msg: msg.RunningScriptMsg{Script: scripts.JsFuntionalMC}},
}

func InitialjsfuntionalModel() menu.Model {

	return menu.InitialModel(jsfuntionalOptions)
}
