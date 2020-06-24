package main

import (
	"fmt"
	"path/filepath"
	"strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// InitGui starts up the whole interface for out program.
func InitGui() {
	// Initialize our new fyne interface application.
	app := app.New()

	// Create the window for our user interface.
	window := app.NewWindow("CNG Medley PDF Parser")

	// Make a vertical box to house all our widgets in.
	box := widget.NewVBox()

	// Fyne doesn't support file chooser yet so we need to enter the filename manually. Should not be a problem for files in the same folder though.
	inputedFile := widget.NewEntry()

	// Add inputedFile to the box.
	box.Append(inputedFile)

	// Create out number input entry.
	inputedNumber := widget.NewEntry()

	// Create the label that acts as a spacer and importing information.
	dataLabel := widget.NewLabel("")

	// Create an array to store recently imported file names.
	var imported []string

	// Create the import button for our file.
	importPDF := widget.NewButton("Importera data från pdf ovan", func() {

		var visitors int

		// Check if filenames are the same.
		if len(imported) != 0 {
			for i := 0; i < len(imported); i++ {
				if imported[i] == inputedFile.Text {
					// Random number that we shouldn't ever get from our function.
					visitors = -5
				}
			}
		}

		extension := ".pdf"

		// If they are not, we run the importer and set text accordingly.
		if visitors != -5 && filepath.Ext(inputedFile.Text) == extension {
			imported = append(imported, inputedFile.Text)
			visitors = Importer(inputedFile.Text)
			dataLabel.SetText("Antal elever under den veckan: " + strconv.Itoa(visitors))
		} else if visitors == -5 && filepath.Ext(inputedFile.Text) == extension {
			dataLabel.SetText("Du har redan importerat filen nyligen: " + inputedFile.Text)
		} else if filepath.Ext(inputedFile.Text) != extension {
			dataLabel.SetText("Vänligen importera en fil med .pdf på slutet.")
		}
	})

	// Add three more widgets to the box.
	box.Append(importPDF)
	box.Append(dataLabel)
	box.Append(inputedNumber)

	// Create the button for showing visitors in the gui.
	displayLessThan := widget.NewButton("Visa antal elever med färre besök än valt nummer ovan", func() {
		// Run the name printer for strings and add the data to the label.
		lessthan, err := strconv.Atoi(inputedNumber.Text)

		if err != nil || lessthan < 2 || lessthan > 10 {
			dataLabel.SetText("Vänligen välj ett nummer mellan (eller lika med) 2 och 10.")
		} else {
			// Resize the window to something that's useable.
			window.Resize(fyne.NewSize(400, 550))

			// Loop through names and add a new label for each name.
			data := ReadDataFromXML()
			for i := 0; i < len(data.Person); i++ {
				if data.Person[i].Visits < lessthan {
					box.Append(widget.NewLabel(fmt.Sprintf("%s: %v besök", data.Person[i].Name, data.Person[i].Visits)))
				}
			}

			// Resize the window a little bit to avoid overlapping data.
			window.Resize(fyne.NewSize(400, 600))
		}
	})

	// Add the last button to the box.
	box.Append(displayLessThan)

	// Make the whole box scrollable.
	scroll := widget.NewScrollContainer(box)

	// Set our content in the application.
	window.SetContent(fyne.NewContainerWithLayout(layout.NewBorderLayout(nil, nil, nil, nil), scroll))

	// Set a default size for the window.
	window.Resize(fyne.NewSize(400, 200))

	// Show all our widgets and initialize out main gui loop.
	window.ShowAndRun()
}
