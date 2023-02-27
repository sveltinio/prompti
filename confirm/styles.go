package confirm

import "github.com/charmbracelet/lipgloss"

// Styles is the struct representing the style configuration options.
type Styles struct {
	Width             int
	BorderColor       lipgloss.AdaptiveColor
	BorderStyle       lipgloss.Border
	MessageStyle      lipgloss.Style
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
	width         = 50
	borderColor   = purple
	borderStyle   = lipgloss.RoundedBorder()
	messageStyle  = lipgloss.NewStyle()
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
			Border(borderStyle).
			BorderForeground(borderColor).
			Margin(1, 0, 0, 0).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).
			Width(width).Align(lipgloss.Center)

	defaultTheme = Styles{
		Width:             width,
		BorderColor:       borderColor,
		BorderStyle:       borderStyle,
		MessageStyle:      messageStyle,
		QuestionStyle:     questionStyle,
		ButtonStyle:       buttonStyle,
		ActiveButtonStyle: activeButtonStyle,
		DialogStyle:       dialogStyle,
	}
)

func setCustomDialogStyles(t *Styles) lipgloss.Style {
	_width := 50
	if !isEmpty(t.Width) {
		_width = t.Width
	}

	_borderStyle := borderStyle
	if !isEmpty((t.BorderStyle)) {
		_borderStyle = t.BorderStyle
	}

	_borderColor := purple
	if !isEmpty((t.BorderColor)) {
		_borderColor = t.BorderColor
	}

	return lipgloss.NewStyle().
		Margin(1, 0, 0, 0).
		Padding(1, 0).
		Border(_borderStyle).
		BorderForeground(_borderColor).
		BorderTop(true).
		BorderLeft(true).
		BorderRight(true).
		BorderBottom(true).
		Width(_width).Align(lipgloss.Center)
}

// DefaultStyles sets the default styles theme.
func DefaultStyles() (s Styles) {
	return defaultTheme
}

func (t *Styles) setDefaults() {
	if isEmpty(t.Width) {
		t.Width = defaultTheme.Width
	}

	if isEmpty(t.BorderColor) {
		t.BorderColor = defaultTheme.BorderColor
	}

	if isEmpty(t.BorderStyle) {
		t.BorderStyle = defaultTheme.BorderStyle
	}

	if isEmpty(t.MessageStyle) {
		t.MessageStyle = defaultTheme.MessageStyle
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
		t.DialogStyle = setCustomDialogStyles(t)
	}
}
