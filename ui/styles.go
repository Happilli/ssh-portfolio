package ui

import "github.com/charmbracelet/lipgloss"

var (
	colorAccent = lipgloss.Color("#7c3aed")
	colorMuted  = lipgloss.Color("#6B7280")
	colorWhite  = lipgloss.Color("#F9FAFB")
	colorGreen  = lipgloss.Color("#10B981")

	StyleTitle     = lipgloss.NewStyle().Foreground(colorAccent).Bold(true).MarginBottom(1)
	StyleSubTitle  = lipgloss.NewStyle().Foreground(colorMuted).Italic(true)
	StyleBody      = lipgloss.NewStyle().Foreground(colorWhite)
	StyleNav       = lipgloss.NewStyle().Foreground(colorMuted)
	StyleNavActive = lipgloss.NewStyle().Foreground(colorAccent).Bold(true)
	StyleTag       = lipgloss.NewStyle().Foreground(colorGreen).PaddingLeft(1).PaddingRight(1).Border(lipgloss.RoundedBorder()).BorderForeground(colorAccent)
	StyleBox       = lipgloss.NewStyle().Border(lipgloss.RoundedBorder()).BorderForeground(colorAccent).Padding(1, 2)
)

func NavBar(active page) string {
	items := []struct {
		key   string
		label string
		p     page
	}{
		{"1", "home", pageHome},
		{"2", "about", pageAbout},
		{"3", "projects", pageProjects},
	}
	var tabs []string
	for _, item := range items {
		label := "[" + item.key + "]" + item.label
		if item.p == active {
			tabs = append(tabs, StyleNavActive.Render(label))
		} else {
			tabs = append(tabs, StyleNav.Render(label))
		}
	}
	nav := lipgloss.JoinHorizontal(lipgloss.Top, tabs[0], "  ", tabs[1], "  ", tabs[2])
	quit := StyleNav.Render("[q] quit")
	return nav + quit
}
