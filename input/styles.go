package input

import "github.com/charmbracelet/lipgloss"

// Styles is the struct representing the style configuration options.
type Styles struct {
	PrefixIcon       string
	PrefixIconColor  lipgloss.AdaptiveColor
	PromptStyle      lipgloss.Style
	TextStyle        lipgloss.Style
	BackgroundStyle  lipgloss.Style
	PlaceholderStyle lipgloss.Style
	CursorStyle      lipgloss.Style
}

// default styles
var defaultTheme = Styles{
	PrefixIcon:       questionMark,
	PrefixIconColor:  purple,
	TextStyle:        noStyle,
	BackgroundStyle:  backgroundStyle,
	PlaceholderStyle: placeholderStyle,
	CursorStyle:      cursorStyle,
}

const (
	// symbols
	promptMark   = "›"
	checkMark    = "✓"
	questionMark = "?"
	cancelMark   = "✘"
)

var (
	// colors
	green  = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"} // Light: green-800, Dark: green-500
	red    = lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"} // Light: red-500, Dark: red-500
	purple = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500
	// text formatting
	bold = lipgloss.NewStyle().Bold(true).Render
	// styles
	noStyle          = lipgloss.NewStyle()
	placeholderStyle = noStyle.Copy().Faint(true)
	backgroundStyle  = noStyle.Copy()
	cursorStyle      = noStyle.Copy()
	prefixIconStyle  = func(color lipgloss.AdaptiveColor) lipgloss.Style {
		return lipgloss.NewStyle().MarginRight(1).Bold(true).Foreground(color)
	}
	questionStyle   = lipgloss.NewStyle().MarginRight(1).Bold(true)
	promptStyle     = lipgloss.NewStyle().Faint(true)
	cancelMarkStyle = lipgloss.NewStyle().MarginRight(1)
	errorStyle      = lipgloss.NewStyle().MarginRight(1).Bold(true).Foreground(red)

	resultPrefixIconStyle = lipgloss.NewStyle().MarginRight(1).Foreground(green)
	resultQuestionStyle   = lipgloss.NewStyle().MarginRight(1)
	resultPromptStyle     = promptStyle.Copy().MarginRight(1)
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
	if isEmpty(t.PromptStyle) {
		t.PromptStyle = defaultTheme.PromptStyle
	}
	if isEmpty(t.TextStyle) {
		t.TextStyle = defaultTheme.TextStyle
	}
	if isEmpty(t.BackgroundStyle) {
		t.BackgroundStyle = defaultTheme.BackgroundStyle
	}
	if isEmpty(t.PlaceholderStyle) {
		t.PlaceholderStyle = defaultTheme.PlaceholderStyle
	}
	if isEmpty(t.CursorStyle) {
		t.CursorStyle = defaultTheme.CursorStyle
	}
}
