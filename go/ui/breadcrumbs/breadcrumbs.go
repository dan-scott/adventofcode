package breadcrumbs

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
)

type Model struct {
	stack    []string
	emptyStr string
}

func (b *Model) Push(item string) {
	b.stack = append(b.stack, item)
}

func (b *Model) Pop() (string, bool) {
	if len(b.stack) == 0 {
		return "", false
	}

	last := b.stack[len(b.stack)-1]
	b.stack = b.stack[:len(b.stack)-1]
	return last, true
}

func (b *Model) View() string {
	doc := strings.Builder{}
	style := lipgloss.NewStyle().BorderStyle(lipgloss.NormalBorder())
	if len(b.stack) == 0 {
		doc.WriteString(" ")
		doc.WriteString(b.emptyStr)
	} else {
		for _, b := range b.stack {
			doc.WriteString(" > ")
			doc.WriteString(b)
		}
	}
	doc.WriteString(" ")
	return fmt.Sprintf("%s\n", style.Render(doc.String()))
}

func New(emptyStr string) *Model {
	return &Model{
		stack:    []string{},
		emptyStr: emptyStr,
	}
}
