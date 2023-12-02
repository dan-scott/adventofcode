package ui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

func newMainMenu() tea.Model {
	return &mainMenu{
		years:    []string{"2023", "2022", "2021", "2020", "2019"},
		selected: 0,
	}
}

type mainMenu struct {
	years    []string
	selected int
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
		}
	}
	m.selected = (m.selected + len(m.years) + mov) % len(m.years)

	return m, nil
}

func (m *mainMenu) View() string {
	doc := strings.Builder{}

	selectedStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#fc8c03"))

	for i, y := range m.years {
		if i == m.selected {
			doc.WriteString(selectedStyle.Render(y))
		} else {
			doc.WriteString(y)
		}
		doc.WriteString("\n")
	}

	return doc.String()
}
