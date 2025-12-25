package main

import (
	// "fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type todo struct {
	text string
	done bool
}

type model struct {
	todos   []todo
	cursor  int
	message string
	selectedItem string
}

func initialModel() model {
	items := []todo{
		{"Open Huh form", false},
		{"Run Gum confirm", false},
		{"Native Bubble Tea confirm", false},
		{"Quit", false},
	}

	return model{
		todos:   items,
		cursor:  0,
		message: "To do list",
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
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "q":
			return m, tea.Quit
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
					if m.selectedItem != ""{	
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
	var s string
	//title
	s += m.message + "\n\n"
	//render todos
	for i, t := range m.todos {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		checbox := "[ ]"
		if t.done {
			checbox = "[x]"
		}
		text := t.text
		if m.cursor == i {
			text = strings.ToUpper(text)
		}
		line := cursor + " " + checbox + " " + text
		s += line + "\n"
	}
	s +=  m.selectedItem
	s += "\n↑/↓ move • space toggle • d delete • q quit"
	return s
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
