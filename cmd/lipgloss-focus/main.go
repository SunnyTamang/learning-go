package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	//"github.com/charmbracelet/lipgloss"

	// CHANGE THIS to your actual module path
	"github.com/SunnyTamang/learningGo/internal/lipgloss-focus/ui"
)

type model struct {
	width  int
	height int

	//scroll  int
	viewport viewport.Model
	lines    []string
	focused  bool
}

func initialModel() model {
	lines := make([]string, 50)
	for i := range lines {
		lines[i] = fmt.Sprintf("Line %02d — this is scrollable content", i+1)
	}

	return model{
		lines:   lines,
		focused: true,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		bodyHeight := ui.Layout{
			Width:  msg.Width,
			Height: msg.Height,
		}.BodyHeight()

		if m.viewport.Width == 0 {
			m.viewport = viewport.New(msg.Width-2, bodyHeight)
			m.viewport.SetContent(strings.Join(m.lines, "\n"))
		} else {
			m.viewport.Width = msg.Width - 2
			m.viewport.Height = bodyHeight
		}

		return m, nil

	case tea.KeyMsg:
		switch msg.String() {

		// case "j", "down":
		// 	m.viewport.LineDown(1)
		// case "k", "up":
		// 	m.viewport.LineUp(1)
		case "tab":
			m.focused = !m.focused

		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.viewport, cmd = m.viewport.Update(msg)

	return m, cmd

}

func scrollStatus(scroll, total int) string {
	if total == 0 {
		return "0 / 0"
	}
	return fmt.Sprintf("Line %d / %d", scroll+1, total)
}

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "loading..."
	}

	// ---- Layout math (same logic as before) ----

	// innerHeight := m.height - 2
	//
	// headerHeight := 1 + ui.HeaderPaddingY()*2
	// footerHeight := 1 + ui.FooterPaddingY()*2
	//
	// bodyHeight := innerHeight -
	// 	headerHeight -
	// 	footerHeight -
	// 	(ui.SeparatorHeight() * 2)
	//
	// if bodyHeight < 1 {
	// 	bodyHeight = 1
	// }

	// ---- Scroll view ----
	
	footerText := fmt.Sprintf(
		"Line %d / %d  •  ↑ ↓ / j k  •  Tab focus  •  q quit",
		m.viewport.YOffset + 1,
		m.viewport.TotalLineCount(),
	)
	

	layout := ui.Layout{
		Width:  m.width,
		Height: m.height,
		Header: "SCROLL DEMO",
		//Body:    scrollView.View(),
		Footer:  footerText,
		Focused: m.focused,
		ShowHeaderShadow: m.viewport.YOffset > 0,
	}

	// scrollView := ui.ScrollView{
	// 	Width:  m.width - 2,
	// 	Height: layout.BodyHeight(),
	// 	Scroll: m.scroll,
	// 	Lines:  m.lines,
	// }
	// layout.Body = scrollView.View()
	layout.Body = m.viewport.View()

	return layout.View()
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
