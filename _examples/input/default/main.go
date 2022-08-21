package main

import (
	"fmt"

	"github.com/sveltinio/prompti/input"
)

func main() {
	questionPrompt := &input.Config{
		Message:     "What's the name of your project?",
		Placeholder: "Please, provide a name for your project",
		ErrorMsg:    "Project name is mandatory",
	}

	result, _ := input.Run(questionPrompt)
	fmt.Println(result)
}
