package layout

import (
	"aten039/scaff/internal/application/javascript"
	"aten039/scaff/internal/application/language"
	"aten039/scaff/internal/components/menu"

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
	Spinner        spinner.Model
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
		Spinner:      s,
	}
}
