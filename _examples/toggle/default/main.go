package main

import (
	"fmt"

	"github.com/sveltinio/prompti/toggle"
)

func main() {
	result, _ := toggle.Run(&toggle.Config{Question: "Continue?"})
	fmt.Println(result)
}
