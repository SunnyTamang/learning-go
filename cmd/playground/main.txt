package main

import (
	"fmt"
	"os"
	"os/exec"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/huh"
)

// ---------- List item ----------

type item string

func (i item) Title() string       { return string(i) }
func (i item) Description() string { return "" }
func (i item) FilterValue() string { return string(i) }

// ---------- Messages ----------

type gumResultMsg bool
type modalResultMsg bool

// ---------- Model ----------

type model struct {
	list        list.Model
	result      string
	showModal   bool
	modalChoice int // 0 = yes, 1 = no
}

// ---------- Init ----------

func initialModel() model {
	items := []list.Item{
		item("Open Huh form"),
		item("Run Gum confirm"),
		item("Native Bubble Tea confirm"),
		item("Quit"),
	}

	l := list.New(items, list.NewDefaultDelegate(), 30, 10)
	l.Title = "Actions"

	return model{
		list:   l,
		result: "Select an action",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// ---------- Gum command ----------

func gumConfirmCmd() tea.Cmd {
	cmd := exec.Command("gum", "confirm", "Do you like Charm?")
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr


	return tea.ExecProcess(cmd, func(err error) tea.Msg {
		return gumResultMsg(err == nil)
	})
}

// ---------- Update ----------

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {

		case "q", "ctrl+c":
			return m, tea.Quit

		case "enter":
			if m.showModal {
				if m.modalChoice == 0 {
					m.result = "Native modal: Yes ❤️"
				} else {
					m.result = "Native modal: No 💔"
				}
				m.showModal = false
				return m, nil
			}

			switch m.list.SelectedItem().(item) {

			case "Open Huh form":
				var name string
				form := huh.NewForm(
					huh.NewGroup(
						huh.NewInput().
							Title("Your name").
							Value(&name),
					),
				)
				form.Run()
				m.result = "Hello " + name

			case "Run Gum confirm":
				return m, gumConfirmCmd()

			case "Native Bubble Tea confirm":
				m.showModal = true
				m.modalChoice = 0

			case "Quit":
				return m, tea.Quit
			}

		case "left", "right":
			if m.showModal {
				if msg.String() == "left" {
					m.modalChoice = 0
				} else {
					m.modalChoice = 1
				}
			}
		}

	case gumResultMsg:
		if msg {
			m.result = "Gum: User likes Charm ❤️"
		} else {
			m.result = "Gum: User said no 💔"
		}

	case modalResultMsg:
		if msg {
			m.result = "Native modal: Yes ❤️"
		} else {
			m.result = "Native modal: No 💔"
		}
	}

	return m, cmd
}

// ---------- View ----------

func (m model) View() string {
	if m.showModal {
		yes := "[ Yes ]"
		no := "[ No ]"

		if m.modalChoice == 0 {
			yes = "> " + yes
		} else {
			no = "> " + no
		}

		return fmt.Sprintf(
			"\nDo you like Bubble Tea?\n\n%s    %s\n\n(← → to choose, Enter to confirm)",
			yes,
			no,
		)
	}

	return fmt.Sprintf(
		"%s\n\nResult:\n%s\n\n(q or Ctrl+C to quit)",
		m.list.View(),
		m.result,
	)
}

// ---------- Main ----------

func main() {
	p := tea.NewProgram(
		initialModel(),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
}
