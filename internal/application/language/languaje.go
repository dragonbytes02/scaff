package language

import (
	"dragonbytes02/scaff/internal/components/menu"
	"dragonbytes02/scaff/internal/msg"
)

var LanguagesOptions = []menu.OptionsInfo{
	{Description: "Go ->", Msg: msg.ChangeToParadigmMsg{}},
	{Description: "Typescript ", Msg: msg.ChangeToParadigmMsg{}},
	{Description: "Python", Msg: msg.ChangeToParadigmMsg{}},
}

func InitialLanguageModel() menu.Model {

	return menu.InitialModel(LanguagesOptions)
}
