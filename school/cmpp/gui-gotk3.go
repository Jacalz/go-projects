package main

import (
	"log"
	"path/filepath"
	"strconv"

	"github.com/gotk3/gotk3/gtk"
)

// Create and initialize out window.
func setupWindow(title string) *gtk.Window {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}

	// Set title and connect it to destroy event.
	win.SetTitle(title)
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	// Set the default size.
	win.SetDefaultSize(400, 230)
	return win
}

func setupButton(label string, onClick func()) *gtk.Button {
	// Create the button.
	btn, err := gtk.ButtonNewWithLabel(label)
	if err != nil {
		log.Fatalln("Unable to create button:", err)
	}

	// Connect the button to the on click function.
	btn.Connect("clicked", onClick)

	return btn
}

func setupEntry() *gtk.Entry {
	// Create the entry.
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}

	return entry
}

// InitGui starts up the whole user interface for us.
func InitGui() {
	// Initialize gtk without arguments.
	gtk.Init(nil)

	// Run setup of window.
	win := setupWindow("Cng Medley PDF Parser")

	// Get all out widgets for the graphical user interface.
	gui := windowWidgets(win)

	win.Add(gui)

	// Show everything and start the main interfac	grid.SetHe loop.
	win.ShowAll()
	gtk.Main()
}

func windowWidgets(win *gtk.Window) *gtk.Widget {

	// Make a new scrolled window for our grid.
	scrolled, _ := gtk.ScrolledWindowNew(nil, nil)

	// Create a grid to house all out little widgets.
	grid, err := gtk.GridNew()
	if err != nil {
		log.Fatal("Unable to create grid:", err)
	}

	// Set the orientation of our spiffy new grid.
	grid.SetOrientation(gtk.ORIENTATION_VERTICAL)

	// Create a new button with the label show.
	show, err := gtk.ButtonNewWithLabel("Visa antal elever med färre besök än valt nummer ovan")
	if err != nil {
		log.Fatalln("Unable to create button:", err)
	}

	// Make a new entry for inputing text.
	entry, err := gtk.EntryNew()
	if err != nil {
		log.Fatal("Unable to create entry:", err)
	}

	// Add a file chooser so the user can select what file to import.
	file, err := gtk.FileChooserButtonNew("Välj pdf", gtk.FILE_CHOOSER_ACTION_OPEN)
	if err != nil {
		log.Fatalln("Unable to create button:", err)
	}

	// Make a button for importing the data from a file.
	importer, err := gtk.ButtonNewWithLabel("Importera pdf")
	if err != nil {
		log.Fatalln("Unable to create button:", err)
	}

	// Create a new label for displaying text.
	label, err := gtk.LabelNew("All data kommer visas här:")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Add a little spacer to separate some widgets in our grid.
	spacer, err := gtk.LabelNew("\n")
	if err != nil {
		log.Fatal("Unable to create label spacer:", err)
	}

	// Connect the show button to printing out names to the label.
	show.Connect("clicked", func() {
		// Get the text and conver to int.
		text, _ := entry.GetText()
		lessthan, err := strconv.Atoi(text)

		// Handle exceptions and error if input is not a number.
		if err != nil || lessthan < 2 || lessthan > 10 {
			label.SetText("Vänligen välj ett nummer mellan (eller lika med) 2 och 10.")
		} else {
			// Resize to fit all the names.
			win.Resize(400, 800)

			// Output names to our label.
			label.SetText(StringLessThan(lessthan))
		}

	})

	var imported []string

	importer.Connect("clicked", func() {

		var visitors int

		// Check if filenames are the same.
		if len(imported) != 0 {
			for i := 0; i < len(imported); i++ {
				if imported[i] == file.GetFilename() {
					// Random number that we shouldn't ever get from our function.
					visitors = -5
				}
			}
		}

		extension := ".pdf"

		// Handle fileextensions and already inputed files.
		if visitors != -5 && filepath.Ext(file.GetFilename()) == extension {
			imported = append(imported, file.GetFilename())
			visitors = Importer(file.GetFilename())
			label.SetText("Antal elever under den veckan: " + strconv.Itoa(visitors))
		} else if visitors == -5 && filepath.Ext(file.GetFilename()) == extension {
			label.SetText("Du har redan importerat filen nyligen: " + filepath.Base(file.GetFilename()))
		} else if filepath.Ext(file.GetFilename()) != extension {
			label.SetText("Vänligen importera en fil med .pdf på slutet.")
		}
	})

	// Set the spacing for our rows and colums.
	grid.SetRowSpacing(4)

	// Add the file chooser and out import button.
	grid.Add(file)
	grid.Add(importer)

	// Add a little spacer to make it look cleaner.
	grid.Add(spacer)

	// Add our text entry and show button,
	grid.Add(entry)
	grid.Add(show)

	// Add the text output label where we display our data.
	grid.Add(label)

	// Tell the entry that it can't expand and that the label can.
	entry.SetHExpand(false)
	label.SetHExpand(true)

	// Add out grid to the scrolled window.
	scrolled.Add(grid)

	return &scrolled.Container.Widget
}
