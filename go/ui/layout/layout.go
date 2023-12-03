package layout

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Model struct {
	width, height int
	Header        Content
	Footer        Content
	Main          Content
}

type Viewable interface {
	View() string
}

type Content interface {
	Viewable
	SetMaxSize(width, height int)
}

func New() *Model {
	return &Model{}
}

func (m *Model) SetMaxSize(width, height int) {
	m.width = width
	m.height = height
	m.Header.SetMaxSize(width, 3)
	m.Footer.SetMaxSize(width, 3)
}

func (m *Model) View() string {
	headerStr := ""
	mainStr := ""
	footerStr := ""
	lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true, false)
	if m.Header != nil {
		headerStr = m.Header.View()
	}

	if m.Footer != nil {
		footerStr = m.Footer.View()
	}
	headerH, _ := lipgloss.Size(headerStr)
	footerH, _ := lipgloss.Size(footerStr)

	maxMainH := m.height - headerH - footerH - 2
	if m.Main != nil {
		mainStr = m.Main.View()
		m.Main.SetMaxSize(maxMainH, m.width)
	}
	mainH, _ := lipgloss.Size(mainStr)
	if mainH < maxMainH {
		mainStr = fmt.Sprintf("%s%s", mainStr, strings.Repeat("\n", maxMainH-mainH))
	}

	return fmt.Sprintf("%s%s%s", headerStr, mainStr, footerStr)
}
