package main

import (
	"fmt"

	"github.com/charmbracelet/bubbles/list"
	"github.com/sveltinio/prompti/choose"
)

func main() {
	foodSelectionPrompt := &choose.Config{
		Title:    "What do you wanna eat tonight?",
		ErrorMsg: "Please, select your meal.",
	}

	entries := []list.Item{
		choose.Item{Name: "pizza", Desc: "It's always pizza time!"},
		choose.Item{Name: "kebab", Desc: "I feel turkish today, kebab!"},
		choose.Item{Name: "carbonara", Desc: "Carbonara, NO cream, please!"},
	}

	result, _ := choose.Run(foodSelectionPrompt, entries)
	fmt.Println(result)
}
