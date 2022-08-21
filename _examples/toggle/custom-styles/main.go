package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/toggle"
)

var (
	cyan   = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}
	green  = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"} // Light: green-800, Dark: green-500
	red    = lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"} // Light: red-500, Dark: red-500
	purple = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500

	myCustomStyle = toggle.Styles{
		PrefixIcon:        "★",
		PrefixIconColor:   red,
		DialogStyle:       lipgloss.NewStyle().Margin(1, 0),
		ButtonStyle:       lipgloss.NewStyle().Bold(true).Foreground(cyan),
		ActiveButtonStyle: lipgloss.NewStyle().Foreground(green),
	}

	toggleConfig = &toggle.Config{
		Question:          "How do you feel?",
		OkButtonLabel:     "I'm super ok",
		CancelButtonLabel: "Next question, please!",
		Divider:           "|",
		Styles:            myCustomStyle,
	}
)

func main() {
	result, _ := toggle.Run(toggleConfig)
	fmt.Println(result)
}
