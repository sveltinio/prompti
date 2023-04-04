// Package choose ...
package choose

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type errMsg error

// Item represents an entry for choose (select) list.
type Item struct {
	Name string
	Desc string
}

// FilterValue returns the current value of the filter.
func (i Item) FilterValue() string { return i.Name }

type itemDelegate struct {
	ItemIcon          string
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(Item)
	if !ok {
		return
	}

	str := fmt.Sprint(d.ItemIcon + whitespace + i.Desc)

	fn := d.ItemStyle.Render
	if index == m.Index() {
		fn = func(s ...string) string {
			return d.SelectedItemStyle.Render(concatStrings("", s...))
		}
	}

	fmt.Fprint(w, fn(str))
}

type model struct {
	list     list.Model
	err      error
	choice   string
	quitting bool
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEscape:
			m.quitting = true
			return m, tea.Quit
		case tea.KeyEnter:
			i, ok := m.list.SelectedItem().(Item)
			if ok {
				m.choice = string(i.Name)
			}
			return m, tea.Quit
		}
	case errMsg:
		m.err = msg
		return m, nil
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return m.choice
	}
	return m.list.View()
}
