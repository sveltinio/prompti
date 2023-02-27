// Package confirm ...
package confirm

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	message           string
	question          string
	okButtonLabel     string
	cancelButtonLabel string
	quitting          bool
	confirmation      bool

	// styles
	styles Styles
}

func (m model) Init() tea.Cmd { return nil }

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc", "q", "n", "N":
			m.confirmation = false
			m.quitting = true
			return m, tea.Quit
		case "left", "h", "ctrl+p", "tab",
			"right", "l", "ctrl+n", "shift+tab":
			m.confirmation = !m.confirmation
		case "enter":
			m.quitting = true
			return m, tea.Quit
		case "y", "Y":
			m.quitting = true
			m.confirmation = true
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m model) View() string {
	if m.quitting {
		return ""
	}

	var aff, neg string

	if m.confirmation {
		aff = m.styles.ActiveButtonStyle.Render(m.okButtonLabel)
		neg = m.styles.ButtonStyle.Render(m.cancelButtonLabel)
	} else {
		aff = m.styles.ButtonStyle.Render(m.okButtonLabel)
		neg = m.styles.ActiveButtonStyle.Render(m.cancelButtonLabel)
	}

	message := lipgloss.NewStyle().Render(m.message)
	question := m.styles.QuestionStyle.Render(m.question)
	buttons := lipgloss.JoinHorizontal(lipgloss.Left, aff, neg)

	var ui string
	if !isEmpty(message) {
		ui = m.styles.DialogStyle.Render(
			lipgloss.JoinVertical(lipgloss.Center, message, "\n", question, buttons))
	} else {
		ui = m.styles.DialogStyle.Render(
			lipgloss.JoinVertical(lipgloss.Center, question, buttons))
	}

	return lipgloss.NewStyle().Render(ui)
}
