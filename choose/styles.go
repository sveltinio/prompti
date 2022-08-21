package choose

import (
	"github.com/charmbracelet/lipgloss"
)

// Styles is the struct representing the style configuration options.
type Styles struct {
	PrefixIcon        string
	PrefixIconColor   lipgloss.AdaptiveColor
	TitleStyle        lipgloss.Style
	TitleBarStyle     lipgloss.Style
	ItemIcon          string
	ItemStyle         lipgloss.Style
	SelectedItemStyle lipgloss.Style
}

const (
	// symbols
	promptMark   = "›"
	checkMark    = "✓"
	whitespace   = " "
	questionMark = "?"
)

var (
	// Colors
	cyan   = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}
	green  = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"} // Light: green-800, Dark: green-500
	red    = lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"} // Light: red-500, Dark: red-500
	purple = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500
	// styles
	noStyle         = lipgloss.NewStyle()
	prefixIconStyle = func(color lipgloss.AdaptiveColor) lipgloss.Style {
		return lipgloss.NewStyle().Bold(true).Foreground(color)
	}
	promptStyle     = lipgloss.NewStyle().MarginRight(1).Faint(true)
	resultIconStyle = lipgloss.NewStyle().MarginRight(1).Foreground(green)
	bold            = lipgloss.NewStyle().Bold(true).Render
	// default theme styles
	defaultTitleStyle        = lipgloss.NewStyle().Bold(true)
	defaultTitleBarStyle     = noStyle.Copy().PaddingBottom(1)
	defaultItemIcon          = lipgloss.NewStyle().Render("▸")
	defaultItemStyle         = lipgloss.NewStyle().PaddingLeft(0)
	defaultSelectedItemStyle = defaultItemStyle.Copy().Foreground(cyan)

	defaultTheme = Styles{
		PrefixIcon:        questionMark,
		PrefixIconColor:   purple,
		TitleStyle:        defaultTitleStyle,
		TitleBarStyle:     defaultTitleBarStyle,
		ItemIcon:          defaultItemIcon,
		ItemStyle:         defaultItemStyle,
		SelectedItemStyle: defaultSelectedItemStyle,
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
	if isEmpty(t.TitleStyle) {
		t.TitleStyle = defaultTheme.TitleStyle
	}
	if isEmpty(t.TitleBarStyle) {
		t.TitleBarStyle = defaultTheme.TitleBarStyle
	}
	if isEmpty(t.ItemIcon) {
		t.ItemIcon = defaultTheme.ItemIcon
	}
	if isEmpty(t.ItemStyle) {
		t.ItemStyle = defaultTheme.ItemStyle
	}
	if isEmpty(t.SelectedItemStyle) {
		t.SelectedItemStyle = defaultTheme.SelectedItemStyle
	}
}
