// package ui
//
// import (
// 	"strings"
//
// 	"github.com/charmbracelet/lipgloss"
// )
//
// const LeftPaneWidth = 24
// const (
// 	headerPaddingY  = 1
// 	footerPaddingY  = 1
// 	separatorHeight = 1
// )
//
// type Layout struct {
// 	Width  int
// 	Height int
//
// 	Header string
// 	Left string
// 	Body   string
// 	Footer string
//
// 	//Focused bool
// 	Focus FocusedSection
//
// 	ShowHeaderShadow bool
//
// 	//LeftTopY int
// }
//
// func (l Layout) BodyHeight() int {
// 	if l.Width == 0 || l.Height == 0 {
// 		return 0
// 	}
//
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
// 	return bodyHeight
// }
//
// func (l Layout) LeftPaneTopY() int {
// 	// Border takes 1 line
// 	// Header takes headerHeight
// 	// Shadow (optional) takes 1
// 	// Separator takes 1
//
// 	headerHeight := 1 + headerPaddingY*2
//
// 	y := 1               // top border
// 	y += headerHeight    // header
//
// 	if l.ShowHeaderShadow {
// 		y += 1
// 	}
//
// 	y += separatorHeight // separator below header
//
// 	return y
// }
//
// func (l Layout) View() string {
// 	if l.Width == 0 || l.Height == 0 {
// 		return "loading..."
// 	}
//
// 	innerWidth := l.Width - 2
// 	leftWidth := LeftPaneWidth
// 	rightWidth := innerWidth - leftWidth
//
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
// 	// Header
// 	headerStyle := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Height(headerHeight).
// 		Padding(headerPaddingY, 2).
// 		Bold(true).
// 		Align(lipgloss.Center)
//
// 	// if !l.Focused {
// 	// 	headerStyle = headerStyle.Foreground(lipgloss.Color("240"))
// 	// }
// 	// if l.Focus != FocusHeader {
// 	// 	headerStyle = headerStyle.Foreground(lipgloss.Color("240"))
// 	// }
//
// 	// if l.Focus != FocusLeft{
// 	// 	headerStyle = headerStyle.Foreground(lipgloss.Color("240"))	
// 	// }
//
// 	header := headerStyle.Render(l.Header)
//
// 	// Header shadow (scroll-aware, layout-owned)
// 	var headerShadow string
// 	if l.ShowHeaderShadow {
// 		headerShadow = lipgloss.NewStyle().
// 			Width(innerWidth).
// 			Foreground(lipgloss.Color("240")).
// 			Render(strings.Repeat("─", innerWidth))
// 	}
//
// 	// Separator
// 	separator := strings.Repeat("─", innerWidth)
//
// 	// Body
// 	// body := lipgloss.NewStyle().
// 	// 	Width(innerWidth).
// 	// 	Height(bodyHeight).
// 	// 	Padding(0, 1).
// 	// 	Render(l.Body)
//
// 	leftStyle := lipgloss.NewStyle().
// 		Width(leftWidth).
// 		Height(bodyHeight).
// 		Padding(1)
//
// 	if l.Focus != FocusLeft{
// 		leftStyle = leftStyle.Foreground(lipgloss.Color("240"))
// 	}
// 	// leftPane:= leftStyle.Render(
// 	// 	"▶ Item 1\n  Item 2\n  Item 3\n  Item 4",
// 	// 	)
// 	leftPane := leftStyle.Render(l.Left)
//
// 	bodyStyle := lipgloss.NewStyle().
// 		//Width(innerWidth).
// 		Width(rightWidth).
// 		Height(bodyHeight).
// 		Padding(0, 1)
//
// 	if l.Focus != FocusBody {
// 		bodyStyle = bodyStyle.Foreground(lipgloss.Color("240"))
// 	}
//
// 	//body := bodyStyle.Render(l.Body)
// 	bodyPane := bodyStyle.Render(l.Body)
// 	bodyRow := lipgloss.JoinHorizontal(
// 		lipgloss.Top,
// 		leftPane,
// 		bodyPane,
// 		)
//
// 	// Footer
// 	footerStyle := lipgloss.NewStyle().
// 		Width(innerWidth).
// 		Height(footerHeight).
// 		Padding(footerPaddingY, 1).
// 		Align(lipgloss.Center)
//
// 	// if !l.Focused {
// 	// 	footerStyle = footerStyle.Foreground(lipgloss.Color("240"))
// 	// }
// 	if l.Focus != FocusFooter {
// 		footerStyle = footerStyle.Foreground(lipgloss.Color("240"))
// 	}
//
// 	footer := footerStyle.Render(l.Footer)
//
// 	content := lipgloss.JoinVertical(
// 		lipgloss.Left,
// 		header,
// 		headerShadow,
// 		//separator,
// 		//body,
// 		bodyRow,
// 		separator,
// 		footer,
// 	)
//
// 	borderColor := lipgloss.Color("240")
// 	// // if !l.Focused {
// 	// // 	borderColor = lipgloss.Color("240")
// 	// // }
// 	// if l.Focus == FocusHeader || l.Focus == FocusBody || l.Focus == FocusFooter {
// 	// 	// active → keep bright border
// 	// } else {
// 	// 	borderColor = lipgloss.Color("240")
// 	// }
// 	switch l.Focus {
// 	//case FocusHeader:
// 	case FocusLeft:
// 		borderColor = lipgloss.Color("69")
// 	case FocusBody:
// 		borderColor = lipgloss.Color("63")
// 	case FocusFooter:
// 		borderColor = lipgloss.Color("141")
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


package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

const (
	LeftPaneWidth = 24

	headerPaddingY  = 1
	footerPaddingY  = 1
	separatorHeight = 1
)

type Layout struct {
	Width  int
	Height int

	Header string
	Left   string
	Body   string
	Footer string

	Focus FocusedSection

	ShowHeaderShadow bool
}

/* ---------- Layout math ---------- */

func headerHeight() int {
	return 1 + headerPaddingY*2
}

func footerHeight() int {
	return 1 + footerPaddingY*2
}

func (l Layout) BodyHeight() int {
	inner := l.Height - 2
	h := inner -
		headerHeight() -
		footerHeight() -
		(separatorHeight * 2)

	if h < 1 {
		return 1
	}
	return h
}

// Absolute Y where left pane content starts (screen coordinates)
func (l Layout) LeftPaneTopY() int {
	y := 1                 // border
	y += headerHeight()    // header
	if l.ShowHeaderShadow {
		y++
	}
	y += separatorHeight  // separator
	y += 1                // left pane top padding (Padding(1))
	return y
}

/* ---------- View ---------- */

func (l Layout) View() string {
	if l.Width == 0 || l.Height == 0 {
		return "loading..."
	}

	innerWidth := l.Width - 2
	rightWidth := innerWidth - LeftPaneWidth

	/* ---------- Header ---------- */

	header := lipgloss.NewStyle().
		Width(innerWidth).
		Height(headerHeight()).
		Padding(headerPaddingY, 2).
		Bold(true).
		Align(lipgloss.Center).
		Render(l.Header)

	var shadow string
	if l.ShowHeaderShadow {
		shadow = lipgloss.NewStyle().
			Width(innerWidth).
			Foreground(lipgloss.Color("240")).
			Render(strings.Repeat("─", innerWidth))
	}

	separator := strings.Repeat("─", innerWidth)

	/* ---------- Left Pane ---------- */

	leftStyle := lipgloss.NewStyle().
		Width(LeftPaneWidth).
		Height(l.BodyHeight()).
		Padding(1)

	if l.Focus != FocusLeft {
		leftStyle = leftStyle.Foreground(lipgloss.Color("240"))
	}

	left := leftStyle.Render(l.Left)

	/* ---------- Body ---------- */

	bodyStyle := lipgloss.NewStyle().
		Width(rightWidth).
		Height(l.BodyHeight()).
		Padding(0, 1)

	if l.Focus != FocusBody {
		bodyStyle = bodyStyle.Foreground(lipgloss.Color("240"))
	}

	body := bodyStyle.Render(l.Body)

	row := lipgloss.JoinHorizontal(lipgloss.Top, left, body)

	/* ---------- Footer ---------- */

	footerStyle := lipgloss.NewStyle().
		Width(innerWidth).
		Height(footerHeight()).
		Padding(footerPaddingY, 1).
		Align(lipgloss.Center)

	if l.Focus != FocusFooter {
		footerStyle = footerStyle.Foreground(lipgloss.Color("240"))
	}

	footer := footerStyle.Render(l.Footer)

	content := lipgloss.JoinVertical(
		lipgloss.Left,
		header,
		shadow,
		row,
		separator,
		footer,
	)

	borderColor := lipgloss.Color("240")
	switch l.Focus {
	case FocusLeft:
		borderColor = lipgloss.Color("69")
	case FocusBody:
		borderColor = lipgloss.Color("63")
	case FocusFooter:
		borderColor = lipgloss.Color("141")
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
