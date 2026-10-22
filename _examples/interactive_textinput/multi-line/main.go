package main

import (
	"github.com/unhealme/pterm"
)

func main() {
	// In multi-line mode, enter inserts a new line and tab submits the input.
	textInput := pterm.DefaultInteractiveTextInput.WithMultiLine()

	result, _ := textInput.Show()

	pterm.Println()
	pterm.Info.Printfln("You answered: %s", result)
}
