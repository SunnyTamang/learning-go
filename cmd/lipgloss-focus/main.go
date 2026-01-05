// package main
//
// import (
// 	"fmt"
// 	"log"
// 	"strings"
//
// 	"github.com/charmbracelet/bubbles/viewport"
// 	tea "github.com/charmbracelet/bubbletea"
// 	//"github.com/charmbracelet/lipgloss"
//
// 	// CHANGE THIS to your actual module path
// 	"github.com/SunnyTamang/learningGo/internal/lipgloss-focus/ui"
// )
//
// type model struct {
// 	width  int
// 	height int
//
// 	//scroll  int
// 	viewport viewport.Model
// 	lines    []string
// 	// focused  bool
// 	focus ui.FocusedSection
//
// 	//left pane state
// 	leftItems []string
// 	leftIndex int
//
// 	leftTopY int
// }
//
// func initialModel() model {
// 	lines := make([]string, 50)
// 	for i := range lines {
// 		lines[i] = fmt.Sprintf("Line %02d — this is scrollable content", i+1)
// 	}
//
// 	return model{
// 		lines: lines,
// 		//focused: true,
// 		focus: ui.FocusBody,
//
// 		leftItems: []string{
// 			"Overview",
// 			"Commits",
// 			"Branches",
// 			"Settings",
// 		},
// 		leftIndex: 0,
// 	}
// }
//
// func (m model) Init() tea.Cmd {
// 	return nil
// }
//
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	var cmd tea.Cmd
// 	switch msg := msg.(type) {
//
// 	case tea.MouseMsg:
// 		// Only left pane handles mouse clicks
// 		if m.focus != ui.FocusLeft {
// 			return m, nil
// 		}
//
// 		// Only react to left button press
// 		if msg.Button != tea.MouseButtonLeft || msg.Action != tea.MouseActionPress {
// 			return m, nil
// 		}
//
// 		// Map mouse Y → left item index
// 		index := msg.Y - m.leftTopY
// 		if index < 0 || index >= len(m.leftItems) {
// 			return m, nil
// 		}
//
// 		m.leftIndex = index
// 		return m, nil
// 	case tea.WindowSizeMsg:
// 		layout := ui.Layout{
// 			Width:  msg.Width,
// 			Height: msg.Height,
// 		}
// 		m.leftTopY = layout.LeftPaneTopY()
// 		m.width = msg.Width
// 		m.height = msg.Height
//
// 		bodyHeight := ui.Layout{
// 			Width:  msg.Width,
// 			Height: msg.Height,
// 		}.BodyHeight()
//
// 		if m.viewport.Width == 0 {
// 			m.viewport = viewport.New(msg.Width-2, bodyHeight)
// 			m.viewport.SetContent(strings.Join(m.lines, "\n"))
// 		} else {
// 			m.viewport.Width = msg.Width - 2
// 			m.viewport.Height = bodyHeight
// 		}
//
// 		return m, nil
//
// 	case tea.KeyMsg:
// 		switch msg.String() {
//
// 		// case "j", "down":
// 		// 	m.viewport.LineDown(1)
// 		// case "k", "up":
// 		// 	m.viewport.LineUp(1)
// 		case "tab":
// 			//m.focused = !m.focused
// 			m.focus = (m.focus + 1) % 3
// 			return m, nil
//
// 		case "q", "ctrl+c":
// 			return m, tea.Quit
// 		}
// 		if m.focus == ui.FocusLeft {
// 			switch msg.String() {
// 			case "j", "down":
// 				if m.leftIndex < len(m.leftItems)-1 {
// 					m.leftIndex++
// 				}
// 			case "k", "up":
// 				if m.leftIndex > 0 {
// 					m.leftIndex--
// 				}
// 			case "enter":
// 				m.viewport.SetContent(
// 					fmt.Sprintf("Selected: %s\n\n%s",
// 						m.leftItems[m.leftIndex],
// 						strings.Join(m.lines, "\n"),
// 					),
// 				)
// 			}
// 			return m, nil
// 		}
// 	}
// 	// m.viewport, cmd = m.viewport.Update(msg)
// 	if m.focus == ui.FocusBody {
// 		m.viewport, cmd = m.viewport.Update(msg)
// 		return m, cmd
// 	}
//
// 	return m, nil
//
// }
//
// // func scrollStatus(scroll, total int) string {
// // 	if total == 0 {
// // 		return "0 / 0"
// // 	}
// // 	return fmt.Sprintf("Line %d / %d", scroll+1, total)
// // }
//
// func (m model) View() string {
// 	var leftBuilder strings.Builder
//
// 	for i, item := range m.leftItems {
// 		if i == m.leftIndex {
// 			leftBuilder.WriteString("▶ " + item + "\n")
// 		} else {
// 			leftBuilder.WriteString("  " + item + "\n")
// 		}
//
// 	}
// 	if m.width == 0 || m.height == 0 {
// 		return "loading..."
// 	}
//
// 	// ---- Layout math (same logic as before) ----
//
// 	// innerHeight := m.height - 2
// 	//
// 	// headerHeight := 1 + ui.HeaderPaddingY()*2
// 	// footerHeight := 1 + ui.FooterPaddingY()*2
// 	//
// 	// bodyHeight := innerHeight -
// 	// 	headerHeight -
// 	// 	footerHeight -
// 	// 	(ui.SeparatorHeight() * 2)
// 	//
// 	// if bodyHeight < 1 {
// 	// 	bodyHeight = 1
// 	// }
//
// 	// ---- Scroll view ----
//
// 	footerText := fmt.Sprintf(
// 		"Line %d / %d  •  ↑ ↓ / j k  •  Tab focus  •  q quit",
// 		m.viewport.YOffset+1,
// 		m.viewport.TotalLineCount(),
// 	)
//
// 	layout := ui.Layout{
// 		Width:  m.width,
// 		Height: m.height,
// 		Header: "SCROLL DEMO",
// 		Left:   leftBuilder.String(),
// 		//Body:    scrollView.View(),
// 		Footer:           footerText,
// 		Focus:            m.focus,
// 		ShowHeaderShadow: m.viewport.YOffset > 0,
// 	}
//
// 	// scrollView := ui.ScrollView{
// 	// 	Width:  m.width - 2,
// 	// 	Height: layout.BodyHeight(),
// 	// 	Scroll: m.scroll,
// 	// 	Lines:  m.lines,
// 	// }
// 	// layout.Body = scrollView.View()
// 	layout.Body = m.viewport.View()
//
// 	return layout.View()
// }
//
// func main() {
// 	f, err := tea.LogToFile("debug.log", "debug")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
//
// 	p := tea.NewProgram(
// 		initialModel(),
// 		tea.WithMouseCellMotion(),
// 	)
// 	if _, err := p.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }


package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"

	"github.com/SunnyTamang/learningGo/internal/lipgloss-focus/ui"
)

type model struct {
	width  int
	height int

	viewport viewport.Model
	lines    []string

	focus ui.FocusedSection

	leftItems []string
	leftIndex int
	leftTopY  int
}

func initialModel() model {
	lines := make([]string, 50)
	for i := range lines {
		lines[i] = fmt.Sprintf("Line %02d — scrollable content", i+1)
	}

	return model{
		lines: lines,
		focus: ui.FocusBody,

		leftItems: []string{
			"Overview",
			"Commits",
			"Branches",
			"Settings",
		},
	}
}

func (m model) Init() tea.Cmd { return nil }

/* ---------- Update ---------- */

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.MouseMsg:
		// Left pane click
		if m.focus == ui.FocusLeft &&
			msg.Button == tea.MouseButtonLeft &&
			msg.Action == tea.MouseActionPress {

			index := msg.Y - m.leftTopY
			if index >= 0 && index < len(m.leftItems) {
				m.leftIndex = index
				return m, nil
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

		bodyH := ui.Layout{
			Width:  msg.Width,
			Height: msg.Height,
		}.BodyHeight()

		if m.viewport.Width == 0 {
			m.viewport = viewport.New(msg.Width-ui.LeftPaneWidth-2, bodyH)
			m.viewport.SetContent(strings.Join(m.lines, "\n"))
		} else {
			m.viewport.Width = msg.Width - ui.LeftPaneWidth - 2
			m.viewport.Height = bodyH
		}
		return m, nil

	case tea.KeyMsg:
		switch msg.String() {
		case "tab":
			m.focus = (m.focus + 1) % ui.FocusCount
			return m, nil

		case "q", "ctrl+c":
			return m, tea.Quit
		}

		if m.focus == ui.FocusLeft {
			switch msg.String() {
			case "j", "down":
				if m.leftIndex < len(m.leftItems)-1 {
					m.leftIndex++
				}
			case "k", "up":
				if m.leftIndex > 0 {
					m.leftIndex--
				}
			case "enter":
				m.viewport.SetContent(
					fmt.Sprintf(
						"Selected: %s\n\n%s",
						m.leftItems[m.leftIndex],
						strings.Join(m.lines, "\n"),
					),
				)
			}
			return m, nil
		}
	}

	// Body input (keyboard + mouse wheel)
	if m.focus == ui.FocusBody {
		m.viewport, cmd = m.viewport.Update(msg)
		return m, cmd
	}

	return m, nil
}

/* ---------- View ---------- */

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "loading..."
	}

	var left strings.Builder
	for i, item := range m.leftItems {
		if i == m.leftIndex {
			left.WriteString("▶ " + item + "\n")
		} else {
			left.WriteString("  " + item + "\n")
		}
	}

	layout := ui.Layout{
		Width:            m.width,
		Height:           m.height,
		Header:           "LIPGLOSS FOCUS DEMO",
		Left:             left.String(),
		Body:             m.viewport.View(),
		Footer:           fmt.Sprintf("Tab: focus • q: quit"),
		Focus:            m.focus,
		ShowHeaderShadow: m.viewport.YOffset > 0,
	}

	m.leftTopY = layout.LeftPaneTopY()

	return layout.View()
}

/* ---------- main ---------- */

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	p := tea.NewProgram(
		initialModel(),
		tea.WithMouseCellMotion(),
	)

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
