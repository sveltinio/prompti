package confirm

import "github.com/charmbracelet/lipgloss"

// Styles is the struct representing the style configuration options.
type Styles struct {
	QuestionStyle     lipgloss.Style
	ButtonStyle       lipgloss.Style
	ActiveButtonStyle lipgloss.Style
	DialogStyle       lipgloss.Style
}

var (
	// Colors
	purple  = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500
	neutral = lipgloss.AdaptiveColor{Light: "737373", Dark: "#a3a3a3"}  // Light: neutral-500, Dark: neutral-400
	amber   = lipgloss.AdaptiveColor{Light: "#fef3c7", Dark: "#fef3c7"} // Light: amber-100, Dark: amber-100

	// Styles
	questionStyle = lipgloss.NewStyle().Bold(true)
	buttonStyle   = lipgloss.NewStyle().
			Foreground(amber).
			Background(neutral).
			Padding(0, 3).
			Margin(1, 1)
	activeButtonStyle = buttonStyle.Copy().
				Foreground(amber).
				Background(purple).
				Underline(true)
	dialogStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(purple).
			Margin(1, 0, 0, 0).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).
			Width(50).Align(lipgloss.Center)

	defaultTheme = Styles{
		QuestionStyle:     questionStyle,
		ButtonStyle:       buttonStyle,
		ActiveButtonStyle: activeButtonStyle,
		DialogStyle:       dialogStyle,
	}
)

// DefaultStyles sets the default styles theme.
func DefaultStyles() (s Styles) {
	return defaultTheme
}

func (t *Styles) setDefaults() {
	if isEmpty(t.QuestionStyle) {
		t.QuestionStyle = defaultTheme.QuestionStyle
	}
	if isEmpty(t.ButtonStyle) {
		t.ButtonStyle = defaultTheme.ButtonStyle
	}
	if isEmpty(t.ActiveButtonStyle) {
		t.ActiveButtonStyle = defaultTheme.ActiveButtonStyle
	}
	if isEmpty(t.DialogStyle) {
		t.DialogStyle = defaultTheme.DialogStyle
	}
}
