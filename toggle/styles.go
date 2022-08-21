package toggle

import "github.com/charmbracelet/lipgloss"

// Styles is the struct representing the style configuration options.
type Styles struct {
	PrefixIcon        string
	PrefixIconColor   lipgloss.AdaptiveColor
	QuestionStyle     lipgloss.Style
	ButtonStyle       lipgloss.Style
	ActiveButtonStyle lipgloss.Style
	DialogStyle       lipgloss.Style
	DividerStyle      lipgloss.Style
}

const (
	questionMark = "?"
	okLabel      = "Yes"
	cancelLabel  = "No"
	cursorLabel  = ">"
	divider      = "/"
)

var (
	// Colors
	cyan   = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}
	muted  = lipgloss.AdaptiveColor{Light: "#D9DCCF", Dark: "#383838"}
	purple = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500
	// Styles
	prefixIconStyle = func(color lipgloss.AdaptiveColor) lipgloss.Style {
		return lipgloss.NewStyle().Bold(true).Foreground(color).PaddingRight(1)
	}
	questionStyle = lipgloss.NewStyle().Bold(true).PaddingRight(1)
	buttonStyle   = lipgloss.NewStyle().
			Foreground(muted).Margin(0)
	activeButtonStyle = buttonStyle.Copy().
				Foreground(cyan).
				Underline(true)
	dialogStyle  = lipgloss.NewStyle()
	dividerStyle = lipgloss.NewStyle().
			Margin(0, 1).
			Foreground(muted)

	defaultTheme = Styles{
		PrefixIcon:        questionMark,
		PrefixIconColor:   purple,
		QuestionStyle:     questionStyle,
		ButtonStyle:       buttonStyle,
		ActiveButtonStyle: activeButtonStyle,
		DialogStyle:       dialogStyle,
		DividerStyle:      dividerStyle,
	}
)

// DefaultStyles sets the default styles theme.
func DefaultStyles() (s Styles) {
	return defaultTheme
}

func (t *Styles) setDefaults() {
	if isEmpty(t.PrefixIcon) {
		t.PrefixIcon = defaultTheme.PrefixIcon
	}
	if isEmpty(t.PrefixIconColor) {
		t.PrefixIconColor = defaultTheme.PrefixIconColor
	}
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
	if isEmpty(t.DividerStyle) {
		t.DividerStyle = defaultTheme.DividerStyle
	}
}
