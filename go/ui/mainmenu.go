package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gitlab.com/danscott/adventofcode/go/ui/breadcrumbs"
	"strings"
)

func newMainMenu() tea.Model {
	return &mainMenu{
		years:    []string{"2023", "2022", "2021", "2020", "2019"},
		selected: 0,
		bcr:      breadcrumbs.New("Select a year..."),
	}
}

type mainMenu struct {
	years         []string
	selected      int
	bcr           *breadcrumbs.Model
	height, width int
	viewport      viewport.Model
}

func (m *mainMenu) Init() tea.Cmd {
	return nil
}

func (m *mainMenu) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	mov := 0
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "up":
			mov = -1
		case "down":
			mov = 1
		case "enter":
			m.bcr.Push(m.years[m.selected])
		case "esc", "backspace":
			m.bcr.Pop()
		}
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
	}

	m.selected = (m.selected + len(m.years) + mov) % len(m.years)

	return m, nil
}

func (m *mainMenu) View() string {
	doc := strings.Builder{}

	doc.WriteString(m.bcr.View())

	selectedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#fc8c03"))

	for i, y := range m.years {
		if i == m.selected {
			doc.WriteString(selectedStyle.Render(y))
		} else {
			doc.WriteString(y)
		}
		doc.WriteString("\n")
	}
	doc.WriteString(fmt.Sprintf("%d - %d", m.width, m.height))

	return doc.String()
}
