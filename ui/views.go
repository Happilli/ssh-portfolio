package ui

import (
	"github.com/charmbracelet/lipgloss"
	"os"
	"strings"
)

// home section
func Homeview(m Model) string {
	ascii, err := os.ReadFile("assets/ascii.txt")
	if err != nil {
		ascii = []byte("RyuZinOh")
	}
	name := StyleTitle.Render(string(ascii))
	slogan := StyleSubTitle.Render("Everything seems like physics, if you don't know magick...")
	hint := StyleNav.Render("\n Use [1-3] to navigate")

	content := lipgloss.JoinVertical(lipgloss.Left, name, slogan, hint)
	return center(m.width, m.height-3, content) + "\n" + NavBar(pageHome)
}

// about section
func AboutView(m Model) string {
	title := StyleTitle.Render("about")
	bio := StyleBody.Render(strings.TrimSpace(`
hey, i'm safal Lama, A Water Enthusiast.
	`))
	content := lipgloss.JoinVertical(lipgloss.Left, title, "", bio, "")

	return padTop(StyleBox.Width(m.width-10).Render(content), 2) + "\n\n" + NavBar(pageAbout)
}

type project struct {
	name, desc, url string
}

var projects = []project{
	{
		name: "edgingIcon",
		desc: "Edges icons at hyprland windows | a hyprland plugin for custom usage",
		url:  "github.com/happilli/edgingIcon",
	},
}

// project section
func ProjectsView(m Model) string {
	title := StyleTitle.Render("projects")
	var cards []string
	for _, p := range projects {
		card := StyleBox.Width(m.width/2 - 6).Render(
			lipgloss.JoinVertical(lipgloss.Left, StyleNavActive.Render(p.name), "",
				StyleBody.Render(p.desc), "",
				StyleSubTitle.Render(p.url),
			),
		)
		cards = append(cards, card)
	}
	row := lipgloss.JoinHorizontal(lipgloss.Top, cards...)
	content := lipgloss.JoinVertical(lipgloss.Left, title, "", row)
	return padTop(content, 2) + "\n\n" + NavBar(pageProjects)
}

func center(w, h int, s string) string {
	return lipgloss.Place(w, h, lipgloss.Center, lipgloss.Center, s)
}

func padTop(s string, n int) string {
	return strings.Repeat("\n", n) + s
}
