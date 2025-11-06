package gotemplates

import (
	"aten039/scaff/internal/components/menu"
	"aten039/scaff/internal/msg"
	"aten039/scaff/internal/scripts"
)

var goOptions = []menu.OptionsInfo{
	{Description: "Clean Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.GoClean}},
	{Description: "Hexagonal Arquitecture", Msg: msg.RunningScriptMsg{Script: scripts.GoHexa}},
	{Description: "Model View Controller", Msg: msg.RunningScriptMsg{Script: scripts.GoMvc}},
	{Description: "Modular Clean", Msg: msg.RunningScriptMsg{Script: scripts.GoMC}},
}

func InitialJsPooModel() menu.Model {

	return menu.InitialModel(goOptions)
}
