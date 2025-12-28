package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	headerPaddingY  = 1
	footerPaddingY  = 1
	separatorHeight = 1
)

type Layout struct {
	Width  int
	Height int

	Header string
	Body   string
	Footer string

	Focused bool

	ShowHeaderShadow bool
}

func (l Layout) BodyHeight() int {
	if l.Width == 0 || l.Height == 0 {
		return 0
	}

	innerHeight := l.Height - 2

	headerHeight := 1 + headerPaddingY*2
	footerHeight := 1 + footerPaddingY*2

	bodyHeight := innerHeight -
		headerHeight -
		footerHeight -
		(separatorHeight * 2)

	if bodyHeight < 1 {
		bodyHeight = 1
	}

	return bodyHeight
}

func (l Layout) View() string {
	if l.Width == 0 || l.Height == 0 {
		return "loading..."
	}

	innerWidth := l.Width - 2
	innerHeight := l.Height - 2

	headerHeight := 1 + headerPaddingY*2
	footerHeight := 1 + footerPaddingY*2

	bodyHeight := innerHeight -
		headerHeight -
		footerHeight -
		(separatorHeight * 2)

	if bodyHeight < 1 {
		bodyHeight = 1
	}

	// Header
	headerStyle := lipgloss.NewStyle().
		Width(innerWidth).
		Height(headerHeight).
		Padding(headerPaddingY, 2).
		Bold(true).
		Align(lipgloss.Center)

	if !l.Focused {
		headerStyle = headerStyle.Foreground(lipgloss.Color("240"))
	}

	header := headerStyle.Render(l.Header)

	// Header shadow (scroll-aware, layout-owned)
	var headerShadow string
	if l.ShowHeaderShadow {
		headerShadow = lipgloss.NewStyle().
			Width(innerWidth).
			Foreground(lipgloss.Color("240")).
			Render(strings.Repeat("─", innerWidth))
	}

	// Separator
	separator := strings.Repeat("─", innerWidth)

	// Body
	body := lipgloss.NewStyle().
		Width(innerWidth).
		Height(bodyHeight).
		Padding(0, 1).
		Render(l.Body)

	// Footer
	footerStyle := lipgloss.NewStyle().
		Width(innerWidth).
		Height(footerHeight).
		Padding(footerPaddingY, 1).
		Align(lipgloss.Center)

	if !l.Focused {
		footerStyle = footerStyle.Foreground(lipgloss.Color("240"))
	}

	footer := footerStyle.Render(l.Footer)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		headerShadow,
		//separator,
		body,
		separator,
		footer,
	)

	borderColor := lipgloss.Color("63")
	if !l.Focused {
		borderColor = lipgloss.Color("240")
	}

	app := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(borderColor).
		Render(content)

	return lipgloss.Place(
		l.Width,
		l.Height,
		lipgloss.Center,
		lipgloss.Center,
		app,
	)
}
