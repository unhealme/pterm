package main

import (
	"github.com/unhealme/pterm"
	"github.com/unhealme/pterm/putils"
)

func main() {
	// Each LettersFromStringWithStyle call gets its own color, so parts of
	// the big text can be highlighted independently.
	pterm.DefaultBigText.WithLetters(
		putils.LettersFromStringWithStyle("P", pterm.FgCyan.ToStyle()),
		putils.LettersFromStringWithStyle("Term", pterm.FgLightMagenta.ToStyle())).
		Render()
}
