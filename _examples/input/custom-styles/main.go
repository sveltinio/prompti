package main

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/input"
)

func main() {
	questionPrompt := &input.Config{
		Message:     "What's the name of your project?",
		Placeholder: "Please, provide a name for your project",
		ErrorMsg:    "Project name is mandatory",
		Styles: input.Styles{
			PrefixIcon:      "*",
			PrefixIconColor: lipgloss.AdaptiveColor{Light: "#ef4444", Dark: "#ef4444"},
			PlaceholderStyle: lipgloss.NewStyle().
				Background(lipgloss.AdaptiveColor{Light: "#3b82f6", Dark: "#3b82f6"}).
				Foreground(lipgloss.AdaptiveColor{Light: "#fde68a", Dark: "#fffbeb"}),
		},
	}

	result, _ := input.Run(questionPrompt)
	fmt.Println(result)
}
