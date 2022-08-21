package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/confirm"
)

var (
	cyan   = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}
	green  = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"} // Light: green-800, Dark: green-500
	red    = lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"} // Light: red-500, Dark: red-500
	purple = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500

	myCustomStyle = confirm.Styles{
		ActiveButtonStyle: lipgloss.NewStyle().Padding(0, 3).
			Margin(1, 1).Background(green).Foreground(purple),
		DialogStyle: lipgloss.NewStyle().
			Border(lipgloss.NormalBorder()).
			BorderTopForeground(purple).
			BorderBottomForeground(green).
			BorderLeftForeground(cyan).
			BorderRightForeground(cyan).
			Margin(1, 0, 1, 0).
			Padding(1, 0).
			BorderTop(true).
			BorderLeft(true).
			BorderRight(true).
			BorderBottom(true).
			Width(50).
			Align(lipgloss.Center),
	}

	confirmConfig = &confirm.Config{
		Question: "Continue?",
		Styles:   myCustomStyle,
	}
)

func main() {
	result, _ := confirm.Run(confirmConfig)
	fmt.Println(result)
}
