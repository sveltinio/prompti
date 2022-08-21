// Package input ...
package input

import (
	"errors"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	textInput    textinput.Model
	message      string
	placeholder  string
	initial      string
	err          error
	errMsg       string
	quitting     bool
	validateFunc ValidateFunc
}

func (m model) Init() tea.Cmd { return textinput.Blink }
func (m model) View() string {
	if isEmpty(m.textInput.Value()) {
		return m.textInput.View()
	}

	if m.err != nil {
		return lipgloss.NewStyle().Render(
			lipgloss.JoinVertical(lipgloss.Left,
				m.textInput.View(),
				errorMessage(m.err.Error())))
	}

	return m.textInput.View()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEscape:
			m.quitting = true
			m.err = errors.New(m.errMsg)
			return m, tea.Quit
		case tea.KeyEnter:
			if !isEmpty(m.initial) && isEmpty(m.textInput.Value()) {
				if m.validateFunc != nil {
					m.err = m.validateFunc(m.initial)
				}
				m.textInput.SetValue(m.initial)
			}

			if m.err == nil || m.validateFunc == nil || m.validateFunc(m.textInput.Value()) == nil {
				return m, tea.Quit
			}
		}
		m.textInput, cmd = m.textInput.Update(msg)
		if m.validateFunc != nil {
			m.err = m.validateFunc(m.textInput.Value())
		}
	case error:
		m.err = msg
		return m, nil
	}
	return m, cmd
}
