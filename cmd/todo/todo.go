package main

import (
	// "fmt"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"log"
	//"strings"
)

var (
	appStyle = lipgloss.NewStyle().Padding(1,2)

	listPaneStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1,2).
		Width(40)
	detailPaneStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("240")).
		Padding(1, 2).
		Width(30)

	addPaneStyle = lipgloss.NewStyle().
		Border(lipgloss.ThickBorder()).
		BorderForeground(lipgloss.Color("212")).
		Padding(1, 2).
		Width(50)
)


var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("205"))

	selectedStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("229"))

	doneStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("42"))

	inputStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("212"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

func (m model) renderListPane() string {
	var s string

	s+= titleStyle.Render("Todos") + "\n\n"
	for i, t := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checkbox := "[ ]"
		if t.done {
			checkbox = "[x]"
		}

		line := cursor + " " + checkbox + " " + t.text

		if t.done {
			line = doneStyle.Render(line)
		}
		if m.cursor == i {
			line = selectedStyle.Render(line)
		}

		s += line + "\n"
	}

	return listPaneStyle.Render(s)
}
func (m model) renderDetailPane() string {
	var s string

	s += titleStyle.Render("Details") + "\n\n"

	if m.selectedItem == "" {
		s += helpStyle.Render("No item selected")
	} else {
		s += m.selectedItem
	}

	return detailPaneStyle.Render(s)
}
type todo struct {
	text string
	done bool
}

type model struct {
	todos        []todo
	cursor       int
	message      string
	selectedItem string

	adding bool
	input  string
}

func initialModel() model {
	items := []todo{
		// {"Open Huh form", false},
		// {"Run Gum confirm", false},
		// {"Native Bubble Tea confirm", false},
		// {"Quit", false},
	}

	return model{
		todos:        items,
		cursor:       0,
		message:      "To do list",
		selectedItem: "",
	}
	// return model{}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.adding {
			switch msg.String() {
			case "esc":
				m.adding = false
				m.input = ""

			case "enter":
				if m.input != "" {
					m.todos = append(m.todos, todo{text: m.input})
					m.cursor = len(m.todos) - 1
				}
				m.adding = false
				m.input = ""

			case "backspace":
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}

			default:
				// add typed characters
				if len(msg.Runes) > 0 {
					m.input += string(msg.Runes)
				}
			}

			return m, nil
		}

		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "a":
			m.adding = true
			m.input = ""
		case "up":
			if len(m.todos) > 0 {
				if m.cursor > 0 {
					m.cursor = (m.cursor - 1)
				} else {
					m.cursor = len(m.todos) - 1
				}
			}
		case "down":
			if len(m.todos) > 0 {
				if m.cursor < len(m.todos)-1 {
					m.cursor = (m.cursor + 1)
				} else {
					m.cursor = 0
				}
			}
		case " ":
			if len(m.todos) > 0 {
				m.todos[m.cursor].done = !m.todos[m.cursor].done
			}
		case "d":
			if len(m.todos) > 0 {
				m.todos = append(m.todos[:m.cursor], m.todos[m.cursor+1:]...)
				if len(m.todos) == 0 {
					m.cursor = 0
					m.selectedItem = ""
				} else {
					if m.selectedItem != "" {
						m.selectedItem = "Selected: " + m.todos[m.cursor].text
					}
				}
			}
		case "enter":
			if len(m.todos) > 0 {
				m.selectedItem = "Selected: " + m.todos[m.cursor].text
			}
		}

	}

	return m, cmd
}

func (m model) View() string {
	// if m.adding {
	// 	// var s string
	// 	// s += titleStyle.Render("Add new todo") + "\n"
	// 	// s += "──────────────\n"
	// 	// s += inputStyle.Render("> "+m.input+"█") + "\n\n"
	// 	// s += helpStyle.Render("(enter to save • esc to cancel)")
	// 	// return s
	// 	return appStyle.Render("Add New Todo") + "\n\n" +
	// 	inputStyle.Render("> "+m.input+"█") + "\n\n"
	// 	helpStyle.Render("(enter to save • esc to cancel)")
	//
	// }
	// var s string
	// //title
	// s += titleStyle.Render(m.message) + "\n\n"
	// //render todos
	// for i, t := range m.todos {
	// 	cursor := " "
	// 	if m.cursor == i {
	// 		cursor = ">"
	// 	}
	// 	checbox := "[ ]"
	// 	if t.done {
	// 		checbox = "[x]"
	// 	}
	// 	text := t.text
	// 	if m.cursor == i {
	// 		text = strings.ToUpper(text)
	// 	}
	// 	line := cursor + " " + checbox + " " + text
	// 	if t.done {
	// 		line = doneStyle.Render(line)
	// 	}
	// 	if m.cursor == i {
	// 		line = selectedStyle.Render(line)
	// 	}
	// 	s += line + "\n"
	// }
	// s += m.selectedItem
	// s += "\n" + helpStyle.Render(
	// 		"↑/↓ move • space toggle • d delete • q quit • a add",
	// )
	// return s

	if m.adding {
		return appStyle.Render(
			addPaneStyle.Render(
				titleStyle.Render("Add Todo") + "\n\n" +
					inputStyle.Render("> " + m.input + "█") + "\n\n" +
					helpStyle.Render("enter save • esc cancel"),
			),
		)
	}

	left := m.renderListPane()
	right := m.renderDetailPane()

	layout := lipgloss.JoinHorizontal(
		lipgloss.Top,
		left,
		right,
	)

	return appStyle.Render(
		layout + "\n\n" +
			helpStyle.Render("↑/↓ move • a add • space toggle • d delete • q quit"),
	)
}

func main() {
	// items := []todo{
	// 	"Get this done",
	// 	"Get that done",
	// }
	// m := New(items)
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %w", err)
	}
	defer f.Close()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
