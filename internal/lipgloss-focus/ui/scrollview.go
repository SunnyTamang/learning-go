package ui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

type ScrollView struct {
	Width  int
	Height int

	Scroll int
	Lines  []string
}

func (s ScrollView) View() string {
	if s.Height <= 0 {
		return ""
	}

	start := s.Scroll
	end := start + s.Height

	if end > len(s.Lines) {
		end = len(s.Lines)
	}
	if start > end {
		start = end
	}

	content := strings.Join(s.Lines[start:end], "\n")

	return lipgloss.NewStyle().
		Width(s.Width).
		Height(s.Height).
		Padding(0, 1).
		Render(content)
}
