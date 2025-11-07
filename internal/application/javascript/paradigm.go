package javascript

import (
	"dragonbytes02/scaff/internal/components/menu"
	"dragonbytes02/scaff/internal/msg"
)

var LanguagesOptions = []menu.OptionsInfo{
	{Description: "object-oriented programming", Msg: msg.ChangeToJsPooMsg{}},
	{Description: "functional programming", Msg: msg.ChangeToJsFuntionalMsg{}},
}

func InitialParadigmModel() menu.Model {

	return menu.InitialModel(LanguagesOptions)
}
