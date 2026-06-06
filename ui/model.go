package ui

import tea "github.com/charmbracelet/bubbletea"

type page int

const (
	pageHome page = iota
	pageAbout
	pageProjects
)

type Model struct {
	width       int
	height      int
	currentPage page
}

func NewModel(w, h int) Model {
	return Model{
		width:       w,
		height:      h,
		currentPage: pageHome,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "1":
			m.currentPage = pageHome
		case "2":
			m.currentPage = pageAbout
		case "3":
			m.currentPage = pageProjects
		}
	}
	return m, nil
}

func (m Model) View() string {
	switch m.currentPage {
	case pageHome:
		return Homeview(m)
	case pageAbout:
		return AboutView(m)
	case pageProjects:
		return ProjectsView(m)
	default:
		return Homeview(m)
	}
}
