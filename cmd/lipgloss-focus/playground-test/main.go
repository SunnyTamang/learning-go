package main

import (
	"fmt"
	"log"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var titleStyle = lipgloss.NewStyle().
	Bold(true).Foreground(lipgloss.Color("205")).
	Padding(1, 2).
	Margin(1, 0)

var containerStyle = lipgloss.NewStyle().
	Border(lipgloss.RoundedBorder()).
	BorderForeground(lipgloss.Color("63")).
	Padding(1, 2).
	Margin(1, 2)

func headerView(width int) string {
	style := lipgloss.NewStyle().
		Width(width).
		Padding(0, 1)
		//Border(lipgloss.NormalBorder())
	return style.Render("Header")
}
func bodyView(width, height int) string {
	style := lipgloss.NewStyle().
		Width(width).
		Height(height).
		Padding(1)
		//Border(lipgloss.NormalBorder())

	return style.Render("body")
}
func footerView(width int) string {
	style := lipgloss.NewStyle().
		Width(width).
		Padding(0, 1)
		//Border(lipgloss.NormalBorder())

	return style.Render("FOOTER")
}

func separator(width int) string {
	return lipgloss.NewStyle().
		Width(width).
		Foreground(lipgloss.Color("26")).
		BorderTop(true).
		BorderStyle(lipgloss.NormalBorder()).
		//Render("")
		Render(strings.Repeat("─", width))
}

type model struct {
	width  int
	height int

	scroll int
	lines  []string
}


func initialModel() model {
	lines := make([]string, 50)
	for i := range lines {
		lines[i] = fmt.Sprintf("Line %02d - this is scrollable content", i+1)
	}
	return model{
		lines: lines,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			if m.scroll < len(m.lines)-1 {
				m.scroll++
			}
		case "k", "up":
			if m.scroll > 0 {
				m.scroll--
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil

}

// func (m model) View() string {
// 	// content := titleStyle.Render("Focus TUI")
// 	// box := containerStyle.Render(content)
// 	// return lipgloss.Place(
// 	// 	m.width,
// 	// 	m.height,
// 	// 	lipgloss.Center,
// 	// 	lipgloss.Center,
// 	// 	box,
// 	// )
// 	// return lipgloss.JoinVertical(
// 	// 	lipgloss.Left,
// 	// 	headerView(m.width),
// 	// 	bodyView(m.width, m.height),
// 	// 	footerView(m.width),
// 	// )
//
// 	if m.width == 0 || m.height == 0 {
// 		return "loading..."
// 	}
//
// 	// Border consumes 2 rows + 2 columns
// 	innerWidth := m.width - 2
// 	innerHeight := m.height - 2
//
// 	headerPaddingY := 1
// 	footerPaddingY := 1
//
// 	headerHeight := 1 + (headerPaddingY * 2)
// 	footerHeight := 1 + (footerPaddingY * 2)
// 	separatorHeight := 1
//
// 	bodyHeight := innerHeight -
// 		headerHeight -
// 		footerHeight -
// 		(separatorHeight * 2)
//
// 	if bodyHeight < 1 {
// 		bodyHeight = 1
// 	}
//
// 	header := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Height(headerHeight).
// 		Padding(headerPaddingY, 2).
// 		Bold(true).
// 		Align(lipgloss.Center).
// 		Render("HEADER")
//
// 	separator := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Foreground(lipgloss.Color("240")).
// 		Render(strings.Repeat("─", innerWidth))
//
// 	body := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Height(bodyHeight).
// 		Padding(1).
// 		Render("Body content goes here")
//
// 	footer := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Height(footerHeight).
// 		Padding(footerPaddingY, 1).
// 		Align(lipgloss.Center).
// 		Foreground(lipgloss.Color("240")).
// 		Render("FOOTER")
//
// 	content := lipgloss.JoinVertical(
// 		lipgloss.Left,
// 		header,
// 		separator,
// 		body,
// 		separator,
// 		footer,
// 	)
//
// 	app := lipgloss.NewStyle().
// 		Border(lipgloss.NormalBorder()).
// 		Render(content)
//
// 	return lipgloss.Place(
// 		m.width,
// 		m.height,
// 		lipgloss.Center,
// 		lipgloss.Center,
// 		app,
// 	)
// }

func (m model) View() string {
	if m.width == 0 || m.height == 0 {
		return "loading..."
	}

	innerWidth := m.width - 2
	innerHeight := m.height - 2

	headerHeight := 3
	footerHeight := 3
	separatorHeight := 1

	bodyHeight := innerHeight -
		headerHeight -
		footerHeight -
		(separatorHeight * 2)

	if bodyHeight < 1 {
		bodyHeight = 1
	}
	start := m.scroll
	end := start + bodyHeight

	if end > len(m.lines) {
		end = len(m.lines)
	}
	bodyContent := strings.Join(m.lines[start:end], "\n")

	body := lipgloss.NewStyle().
		Width(innerWidth).
		Height(bodyHeight).
		Padding(0, 1).
		Render(bodyContent)

	header := lipgloss.NewStyle().
		Width(innerWidth).
		Height(headerHeight).
		Padding(1, 2).
		Bold(true).
		Align(lipgloss.Center).
		Render("SCROLL DEMO")

	separator := strings.Repeat("─", innerWidth)

	footer := lipgloss.NewStyle().
		Width(innerWidth).
		Height(footerHeight).
		Padding(1).
		Align(lipgloss.Center).
		Foreground(lipgloss.Color("240")).
		Render("↑ ↓ / j k  •  q to quit")

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		separator,
		body,
		separator,
		footer,
	)

	app := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		Render(content)

	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		app,
	)
}

func main() {
	f, err := tea.LogToFile("debug.log", "debug")
	if err != nil {
		log.Fatalf("err: %w", err)
	}
	defer f.Close()
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}

}

// latest working code, above is old working code
// package main
//
// import (
// 	"fmt"
// 	"log"
// 	"strings"
//
// 	tea "github.com/charmbracelet/bubbletea"
// 	"github.com/charmbracelet/lipgloss"
// 	"github.com/SunnyTamang/internal/ui/lipgloss-focus"
// )
//
// const (
// 	headerPaddingY  = 1
// 	footerPaddingY  = 1
// 	separatorHeight = 1
// )
//
// /* ──────────────────────
//    Layout helper
// ────────────────────── */
//
// type Layout struct {
// 	Width  int
// 	Height int
//
// 	Header string
// 	Body   string
// 	Footer string
//
// 	Focused bool
// }
//
// func (l Layout) View() string {
// 	if l.Width == 0 || l.Height == 0 {
// 		return "loading..."
// 	}
//
// 	// Outer border consumes 2 rows + 2 columns
// 	innerWidth := l.Width - 2
// 	innerHeight := l.Height - 2
//
// 	headerHeight := 1 + headerPaddingY*2
// 	footerHeight := 1 + footerPaddingY*2
//
// 	bodyHeight := innerHeight -
// 		headerHeight -
// 		footerHeight -
// 		(separatorHeight * 2)
//
// 	if bodyHeight < 1 {
// 		bodyHeight = 1
// 	}
//
// 	headerStyle := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Height(headerHeight).
// 		Padding(headerPaddingY, 2).
// 		Bold(true).
// 		Align(lipgloss.Center)
// 		// Foreground(
// 		// lipgloss.Color(
// 		// 	map[bool]string{true: "205", false: "240"}[true],
// 		// 	),
// 		// ).
// 		// Align(lipgloss.Center).
// 		// Render(l.Header)
// 		if !l.Focused {
// 	headerStyle = headerStyle.Foreground(lipgloss.Color("240"))
// }
// 	header:= headerStyle.Render(l.Header)
//
// 	separator := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Foreground(lipgloss.Color("240")).
// 		Render(strings.Repeat("─", innerWidth))
//
// 	body := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Height(bodyHeight).
// 		Padding(0, 1).
// 		Render(l.Body)
//
// 	footerStyle := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Height(footerHeight).
// 		Padding(footerPaddingY, 1).
// 		Align(lipgloss.Center)
// 		// Foreground(lipgloss.Color("240")).
// 		// Render(l.Footer)
// 		if !l.Focused {
// 	footerStyle = footerStyle.Foreground(lipgloss.Color("240"))
// }
//
// footer := footerStyle.Render(l.Footer)
//
// 	content := lipgloss.JoinVertical(
// 		lipgloss.Left,
// 		header,
// 		separator,
// 		body,
// 		separator,
// 		footer,
// 	)
//
// 	borderColor := lipgloss.Color("63")
// 	if !l.Focused {
// 		borderColor = lipgloss.Color("240")
// 	}
//
// 	app := lipgloss.NewStyle().
// 		Border(lipgloss.NormalBorder()).
// 		BorderForeground(borderColor).
// 		Render(content)
//
// 	return lipgloss.Place(
// 		l.Width,
// 		l.Height,
// 		lipgloss.Center,
// 		lipgloss.Center,
// 		app,
// 	)
// }
//
// type ScrollView struct {
// 	Width int
// 	Height int
//
// 	Scroll int
// 	Lines []string
// }
//
// func (s ScrollView) View() string{
// 	if s.Height <= 0 {
// 		return ""
// 	}
//
// 	start:=s.Scroll
// 	end:= start + s.Height
//
// 	if end > len(s.Lines){
// 		end = len(s.Lines)
// 	}
// 	if start > end {
// 		start = end
// 	}
//
// 	content := strings.Join(s.Lines[start:end], "\n")
// 	return lipgloss.NewStyle().
// 	Width(s.Width).
// 	Height(s.Height).
// 	Padding(0, 1).
// 	Render(content)
// }
//
// func scrollStatus(scroll, total int) string{
// 	if total == 0{
// 		return "0 / 0"
// 	}
// 	return fmt.Sprintf("Line %d / %d", scroll + 1, total)
// }
//
// /* ──────────────────────
//    Bubble Tea model
// ────────────────────── */
//
// type model struct {
// 	width  int
// 	height int
//
// 	scroll int
// 	lines  []string
//
// 	focused bool
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
// 		focused: true,
// 	}
// }
//
// func (m model) Init() tea.Cmd {
// 	return nil
// }
//
// // /* Extract visible body content based on scroll */
// // func (m model) bodyContent(bodyHeight int) string {
// // 	start := m.scroll
// // 	end := start + bodyHeight
// //
// // 	if end > len(m.lines) {
// // 		end = len(m.lines)
// // 	}
// // 	if start > end {
// // 		start = end
// // 	}
// //
// // 	return strings.Join(m.lines[start:end], "\n")
// // }
//
// func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
// 	switch msg := msg.(type) {
//
// 	case tea.WindowSizeMsg:
// 		m.width = msg.Width
// 		m.height = msg.Height
//
// 	case tea.KeyMsg:
// 		switch msg.String() {
// 		case "tab":
// 		m.focused = !m.focused
// 		case "j", "down":
// 			if m.scroll < len(m.lines)-1 {
// 				m.scroll++
// 			}
// 		case "k", "up":
// 			if m.scroll > 0 {
// 				m.scroll--
// 			}
// 		case "q", "ctrl+c":
// 			return m, tea.Quit
// 		}
// 	}
//
// 	return m, nil
// }
//
// func (m model) View() string {
// 	if m.width == 0 || m.height == 0 {
// 		return "loading..."
// 	}
//
// 	// Compute body height here (same math as Layout)
// 	innerHeight := m.height - 2
// 	headerHeight := 1 + headerPaddingY*2
// 	footerHeight := 1 + footerPaddingY*2
//
// 	bodyHeight := innerHeight -
// 		headerHeight -
// 		footerHeight -
// 		(separatorHeight * 2)
//
// 	if bodyHeight < 1 {
// 		bodyHeight = 1
// 	}
// 	// ScrollView:= ScrollView{
// 	// 	Width: m.width - 2,
// 	// 	Height: bodyHeight,
// 	// 	Scroll: m.scroll,
// 	// 	Lines: m.lines,
// 	// }
// 	scrollView := ui.ScrollView{
// 	Width:  m.width - 2,
// 	Height: bodyHeight,
// 	Scroll: m.scroll,
// 	Lines:  m.lines,
// }
// 	footerText := fmt.Sprintf(
// 		"%s  •  ↑ ↓ / j k  •  q to quit",
// 	scrollStatus(m.scroll, len(m.lines)), 
// 		)
// 	layout := Layout{
// 		Width:  m.width,
// 		Height: m.height,
// 		Header: "SCROLL DEMO from playground",
// 		// Body:   m.bodyContent(bodyHeight),
// 		Body: ScrollView.View(),
// 		Footer: footerText,
// 		Focused: m.focused,
// 	}
//
// 	return layout.View()
// }
// /* ──────────────────────
//    main
// ────────────────────── */
//
// func main() {
// 	f, err := tea.LogToFile("debug.log", "debug")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer f.Close()
//
// 	p := tea.NewProgram(initialModel())
// 	if _, err := p.Run(); err != nil {
// 		log.Fatal(err)
// 	}
// }
//
