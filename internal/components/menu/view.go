package menu

import "fmt"

func (m Model) View() string {
	// The header
	header := HeaderStyle.Render("SCAFF")
	header += " pragmatic architectures"
	// Iterate over our choices
	body := TitleStyle.Render(m.title)
	for i, choice := range m.choices {

		cursor := " "
		if m.cursor == i {
			cursor = CursorStyle.Render("> ")
		}

		body += BodyStyle.Render(fmt.Sprintf("\n%s%s", cursor /*checked,*/, choice.Color.Render(choice.Description)))
	}

	// The footer
	footer := FooterStyle.Render(fmt.Sprintf("Press %s to Exit\n", ColorExit.Render("Q")))

	// Send the UI for rendering
	return fmt.Sprintf("\n%s \n%s \n%s", header, ContainerStyle.Render(body), footer)
}
