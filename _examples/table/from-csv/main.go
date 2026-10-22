package main

import (
	"github.com/unhealme/pterm"
	"github.com/unhealme/pterm/putils"
)

func main() {
	csv := `Firstname,Lastname,Email
Paul,Dean,paul@example.com
Callie,Mckay,callie@example.com
Libby,Camacho,libby@example.com`

	// TableDataFromCSV converts raw CSV into pterm.TableData. The CSV header
	// line ends up as the first row, so WithHasHeader renders it as such.
	pterm.DefaultTable.WithHasHeader().WithData(putils.TableDataFromCSV(csv)).Render()
}
