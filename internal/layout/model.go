package layout

import (
	Go "dragonbytes02/scaff/internal/application/go"
	"dragonbytes02/scaff/internal/application/javascript"
	"dragonbytes02/scaff/internal/application/language"
	"dragonbytes02/scaff/internal/components/menu"

	"github.com/charmbracelet/bubbles/spinner"
)

type AppState int

const (
	LanguageModel AppState = iota
	ParadigmModel
	JsPoo
	JsFuntional
	RunningScript
	ScriptError
	ScriptFinished
	GoOptions
)

type AppModel struct {
	name_project   string
	state          AppState
	ScriptError    string
	ScriptFinished string
	Language       menu.Model
	Paradigm       menu.Model
	JsPoo          menu.Model
	JsFuntional    menu.Model
	GoOptions      menu.Model

	Spinner spinner.Model
}

func InitialApp(name string) AppModel {
	s := spinner.New()
	s.Spinner = spinner.Dot
	// s.Style = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))

	return AppModel{
		name_project: name,
		state:        0,
		Language:     language.InitialLanguageModel(),
		Paradigm:     javascript.InitialParadigmModel(),
		JsPoo:        javascript.InitialJsPooModel(),
		JsFuntional:  javascript.InitialjsfuntionalModel(),
		GoOptions:    Go.InitialGoOptionModel(),
		Spinner:      s,
	}
}
