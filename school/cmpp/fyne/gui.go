package main

import (
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
)

// InitGui starts up the whole interface for out program.
func InitGui() {
	// Initialize our new fyne interface application.
	app := app.New()

	// Create the window for our user interface.
	window := app.NewWindow("CNG Medley PDF Parser")

	// Fyne doesn't support file chooser yet so we need to enter the filename manually. Should not be a problem for files in the same folder though.
	inputedFile := widget.NewEntry()

	// Create out number input entry.
	inputedNumber := widget.NewEntry()

	// Create the label where we should be printing our information to.
	dataLabel := widget.NewLabel("")
	dataLabel.Resize(fyne.NewSize(400, 600))

	spacer := widget.NewLabel("")

	/*
		scrollable := widget.NewScrollContainer(dataLabel)
		scrollable.Resize(fyne.NewSize(400, 600))
	*/

	// Create the import button for our file.
	importPDF := widget.NewButton("Importera", func() {
		visitors := Importer(inputedFile.Text)
		dataLabel.SetText("Antal elever under den veckan: " + strconv.Itoa(visitors))
	})

	// Create the button for showing visitors in the gui.
	displayLessThan := widget.NewButton("Visa antal elever med färre besök än valt nummer ovan", func() {
		window.Resize(fyne.NewSize(400, 600))
		text := inputedNumber.Text
		names := StringLessThan(CheckNumber(text))
		dataLabel.SetText(names)
	})

	// Set our content in the application.
	window.SetContent(widget.NewVBox(
		inputedFile,
		importPDF,
		spacer,
		inputedNumber,
		displayLessThan,
		dataLabel,
	))

	// We are not actually resizing it here, we set the deafult size.
	window.Resize(fyne.NewSize(400, 250))

	// Show all our widgets and initialize out main gui loop.
	window.ShowAndRun()
}