package main

import "github.com/unhealme/pterm"

func main() {
	// The box sizes itself to fit whatever it prints.
	pterm.DefaultBox.Println("Hello, World!")
}
