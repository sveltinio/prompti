package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/lipgloss"
	"github.com/sveltinio/prompti/progressbar"
)

func main() {
	fruits := []string{
		"apple",
		"banana",
		"orange",
		"grapes",
		"mellon",
		"strawberry",
		"mango",
		"lemon",
		"apricot",
		"peach",
		"papaya",
		"kiwi",
		"pear",
		"guava",
		"almond",
		"coconut",
		"blackberry",
		"cherry",
		"grapes",
	}

	pbConfig := &progressbar.Config{
		Items:         fruits,
		OnProgressMsg: "Eating:",
		Styles: progressbar.Styles{
			CurrentItemStyle: lipgloss.NewStyle().Foreground(lipgloss.Color("411")),
			ShowLabel:        true,
			GradientFrom:     "#FF7CCB",
			GradientTo:       "#FDFF8C",
		}}

	if _, err := progressbar.Run(pbConfig); err != nil {
		fmt.Println("error running program:", err)
		os.Exit(1)
	}
}
