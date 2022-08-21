package choose

import "github.com/charmbracelet/lipgloss"

var (
	// messages
	titleMessage = func(icon string, iconColor lipgloss.AdaptiveColor, titleStyle lipgloss.Style, title string) string {
		return lipgloss.NewStyle().Render(
			lipgloss.JoinHorizontal(lipgloss.Center,
				prefixIconStyle(iconColor).Render(icon),
				titleStyle.Render(whitespace),
				titleStyle.Render(title),
				titleStyle.Render(whitespace),
				titleStyle.Render(promptMark)))
	}

	resultMessage = func(text, value string) string {
		return lipgloss.NewStyle().Render(
			lipgloss.JoinHorizontal(lipgloss.Center,
				resultIconStyle.Render(checkMark),
				lipgloss.NewStyle().MarginRight(1).Render(text),
				promptStyle.Render(promptMark),
				bold(value)))
	}
	errorMessage = lipgloss.NewStyle().Foreground(red).Bold(true).Render
)
