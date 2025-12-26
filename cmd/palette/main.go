package main

import (
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type Command struct {
	Name string
	Run  func(*model)
}

type history struct {
	past   []model
	future []model
}

type model struct {
	items    []string
	cursor   int
	selected string

	adding bool
	input  string

	paletteOpen  bool
	query        string
	commands     []Command
	paletteIndex int

	history history
}

func (m model) snapshot() model {
	return model{
		items:    append([]string{}, m.items...),
		cursor:   m.cursor,
		selected: m.selected,
	}
}

func (m *model) pushHistory() {
	m.history.past = append(m.history.past, m.snapshot())
	m.history.future = nil //clear redo stack
}

func (m *model) undo() {
	if len(m.history.past) == 0 {
		return
	}
	prev := m.history.past[len(m.history.past)-1]
	m.history.past = m.history.past[:len(m.history.past)-1]

	m.history.future = append(m.history.future, m.snapshot())
	m.items = prev.items
	m.cursor = prev.cursor
	m.selected = prev.selected
}

func (m *model) redo() {
	if len(m.history.future) == 0 {
		return
	}
	next := m.history.future[len(m.history.future)-1]
	m.history.future = m.history.future[:len(m.history.future)-1]

	m.history.past = append(m.history.past, m.snapshot())
	m.items = next.items
	m.cursor = next.cursor
	m.selected = next.selected
}

func (m model) filteredCommands() []Command {
	if m.query == "" {
		return m.commands
	}
	var result []Command
	for _, c := range m.commands {
		if strings.Contains(
			strings.ToLower(c.Name),
			strings.ToLower(m.query),
		) {
			result = append(result, c)
		}
	}
	return result
}

func initialModel() model {
	listitems := []string{
		"This is number 1",
		"This is number 2",
	}
	listCommands := []Command{
		{
			Name: "Add todo",
			Run: func(m *model) {
				m.adding = true
				m.input = ""
			},
		},
		{
			Name: "Clear Selection",
			Run: func(m *model) {
				m.pushHistory()
				m.selected = ""
			},
		},
	}
	return model{
		items:    listitems,
		cursor:   0,
		selected: "",
		commands: listCommands,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if m.paletteOpen {
			cmds := m.filteredCommands()
			m.selected = ""
			switch msg.String() {
			case "esc":
				m.paletteOpen = false
				m.query = ""
				m.paletteIndex = 0
			case "enter":
				if len(cmds) > 0 {
					cmds[m.paletteIndex].Run(&m)
				}
				m.paletteOpen = false
				m.query = ""
				m.paletteIndex = 0
			case "backspace":
				if len(m.query) > 0 {
					m.query = m.query[:len(m.query)-1]
				}
				cmds = m.filteredCommands()
				if m.paletteIndex >= len(cmds) {
					m.paletteIndex = 0
				}
			case "up":
				if len(cmds) > 0 {
					if m.paletteIndex > 0 {
						m.paletteIndex--
					} else {
						m.paletteIndex = len(cmds) - 1
					}
				}
			case "down":
				if len(cmds) > 0 {
					if m.paletteIndex < len(cmds)-1 {
						m.paletteIndex++
					} else {
						m.paletteIndex = 0
					}
				}
			default:
				if len(msg.Runes) > 0 {
					m.query += string(msg.Runes)
					cmds = m.filteredCommands()
					if m.paletteIndex >= len(cmds) {
						m.paletteIndex = 0
					}
				}
			}
			return m, nil
		} else if m.adding {
			m.selected = ""
			switch msg.String() {
			case "esc":
				m.adding = false
				m.input = ""
			case "enter":
				if m.input != "" {
					m.pushHistory()
					m.items = append(m.items, m.input)
					m.cursor = len(m.items) - 1

				}
				m.adding = false
				m.input = ""

			case "backspace":
				if len(m.input) > 0 {
					m.input = m.input[:len(m.input)-1]
				}
			default:
				if len(msg.Runes) > 0 {
					m.input += string(msg.Runes)
				}
			}
			return m, nil
		} else {
			m.selected = ""
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "/":
				m.paletteOpen = true
				m.query = ""
				m.paletteIndex = 0
			case "a":
				m.adding = true
				m.input = ""
			case "up":
				if len(m.items) > 0 {
					if m.cursor > 0 {
						m.cursor--
					} else {
						m.cursor = len(m.items) - 1
					}
				}

			case "down":
				if len(m.items) > 0 {
					if m.cursor < len(m.items)-1 {
						m.cursor++
					} else {
						m.cursor = 0
					}
				}
			case "enter":
				if len(m.items) > 0 {
					m.selected = "Selected: " + m.items[m.cursor]
				}
			case "u":
				m.undo()

			case "r":
				m.redo()
			}
		}
	}
	return m, cmd
}

func (m model) View() string {
	if m.adding {
		var s string
		s += "Adding todo mode\n\n"
		s += "Add a new item" + "\n"
		s += "---------------\n\n"
		s += "> " + m.input + "_" + "\n\n"
		s += "(enter to save • esc to cancel)"
		return s
	}
	if m.paletteOpen {
		var s string
		cmds := m.filteredCommands()
		s += "Palette mode\n\n"
		s += "/ " + m.query + "_" + "\n"
		s += "--------------------------\n"
		for i, cmd := range cmds {
			palleteindex := " "
			if m.paletteIndex == i {
				palleteindex = ">"
			}
			command_lines := palleteindex + " " + cmd.Name
			s += command_lines + "\n"
		}
		s += "\n(enter to save • esc to cancel)"
		return s
	}
	var s string
	s += "Normal mode\n\n"
	s += "List of Items\n\n"
	for i, t := range m.items {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}
		line := cursor + " " + t
		s += line + "\n"
	}
	s += "\n" + m.selected
	s += "\n↑/↓ move • a add • / palette • u undo • r redo • q quit"
	return s
}

func main() {
	f, err := tea.LogToFile("palette_debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %w", err)
	}
	defer f.Close()
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
