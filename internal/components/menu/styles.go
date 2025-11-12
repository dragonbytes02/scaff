package menu

import "github.com/charmbracelet/lipgloss"

var HeaderStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#8D44F2")).AlignHorizontal(lipgloss.Center)

var ContainerStyle = lipgloss.NewStyle().Width(40)

var BodyStyle = lipgloss.NewStyle().
	PaddingLeft(4).Align(lipgloss.Left).Padding().PaddingTop(1)

var ColorExit = lipgloss.NewStyle().Foreground(lipgloss.Color("#F28322"))

var CursorStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color("#FFFFFF"))

var TitleStyle = lipgloss.NewStyle().Bold(true).Width(40).Align(lipgloss.Center).BorderStyle(lipgloss.NormalBorder()).BorderBottom(true).BorderTop(true).BorderForeground(lipgloss.Color("#0554F2"))

var FooterStyle = lipgloss.NewStyle().Bold(true).Align(lipgloss.Left).MarginTop(1).BorderStyle(lipgloss.NormalBorder()).BorderTop(true).BorderForeground(lipgloss.Color("#0554F2")).Width(40)
