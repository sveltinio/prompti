package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/choose"
)

var (
	amber  = lipgloss.AdaptiveColor{Light: "#f59e0b", Dark: "#fbbf24"} // Light: amber-500, Dark: amber-400
	purple = lipgloss.AdaptiveColor{Light: "#7e22ce", Dark: "#a855f7"} // Light: purple-700, Dark: purple-500
	green  = lipgloss.AdaptiveColor{Light: "#166534", Dark: "#22c55e"} // Light: green-800, Dark: green-500

	myCustomStyle = choose.Styles{
		PrefixIcon:        "★",
		TitleStyle:        lipgloss.NewStyle().Background(green).Foreground(purple).Padding(0, 1),
		TitleBarStyle:     lipgloss.NewStyle(),
		ItemIcon:          "#",
		ItemStyle:         lipgloss.NewStyle().Foreground(amber),
		SelectedItemStyle: lipgloss.NewStyle().Foreground(purple),
	}
)

func main() {
	foodSelectionPrompt := &choose.Config{
		Title:    "What do you wanna eat tonight?",
		ErrorMsg: "Please, select your meal.",
		ShowHelp: true,
		Styles:   myCustomStyle,
	}

	entries := []list.Item{
		choose.Item{Name: "pizza", Desc: "It's always pizza time!"},
		choose.Item{Name: "kebab", Desc: "I feel turkish today, kebab!"},
		choose.Item{Name: "carbonara", Desc: "Carbonara, NO cream, please!"},
	}

	result, _ := choose.Run(foodSelectionPrompt, entries)
	fmt.Println(result)
}
