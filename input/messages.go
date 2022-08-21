package input

import "github.com/charmbracelet/lipgloss"

var (
	// messages
	promptMessage = func(question string, icon string, iconColor lipgloss.AdaptiveColor) string {
		return lipgloss.NewStyle().MarginTop(1).Render(lipgloss.JoinHorizontal(lipgloss.Center,
			prefixIconStyle(iconColor).Render(icon),
			questionStyle.Render(question),
			promptStyle.Render(promptMark)))
	}
	resultMessage = func(question, value string) string {
		return lipgloss.NewStyle().Render(lipgloss.JoinHorizontal(lipgloss.Center,
			resultPrefixIconStyle.Render(checkMark),
			resultQuestionStyle.Render(question),
			resultPromptStyle.Render(promptMark),
			bold(value)))
	}
	errorMessage = func(s string) string {
		return errorStyle.Render(
			lipgloss.JoinHorizontal(lipgloss.Center,
				cancelMarkStyle.Render(cancelMark), s))
	}
)
