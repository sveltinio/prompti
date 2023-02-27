package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/confirm"
)

var (
	cyan  = lipgloss.AdaptiveColor{Light: "#4f46e5", Dark: "#c7d2fe"}
	green = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"} // Light: green-800, Dark:

	infoText = `Lorem ipsum dolor sit amet,
consectetur adipiscing elit %s...`

	Green   = lipgloss.NewStyle().Foreground(green).Render
	message = fmt.Sprintf(infoText, Green("elit"))

	myCustomStyle = confirm.Styles{
		Width:       60,
		BorderColor: cyan,
	}

	confirmConfig = &confirm.Config{
		Message:  message,
		Question: "Continue?",
		Styles:   myCustomStyle,
	}
)

func main() {
	result, _ := confirm.Run(confirmConfig)
	fmt.Println(result)
}
