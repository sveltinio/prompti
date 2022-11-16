package progressbar

import "github.com/charmbracelet/lipgloss"

// Styles is the struct representing the style configuration options.
type Styles struct {
	ShowLabel        bool
	CurrentItemStyle lipgloss.Style
	GradientFrom     string
	GradientTo       string
}
