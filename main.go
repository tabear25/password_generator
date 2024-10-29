package main

import (
	"log"

	"fyne.io/fyne/v2/app"

	"github.com/yourusername/password_generator/ui"
)

func main() {
	myApp := app.New()
	myWindow := ui.InitUI(myApp)
	if myWindow == nil {
		log.Fatal("Failed to initialize the UI")
	}
	myWindow.ShowAndRun()
}
