package main

import (
	"regexp"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("SEABorne")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter regex pattern to look for")

	button := widget.NewButtonWithIcon("Run format", theme.ContentCopyIcon(), func() {
		pattern, err := regexp.Compile(input.Text)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()

		result := pattern.FindAllString(clipboard.Content(), -1)
		if len(result) == 0 {
			dialog.ShowInformation("No results", "Couldn't find anything with that expression.", w)
			return
		}

		formated := ""
		for i, v := range result {
			if i != len(result)-1 {
				formated += v + "\n"
			} else {
				formated += v
			}
		}

		clipboard.SetContent(formated)
	})

	w.Resize(fyne.NewSize(400, 200))

	w.SetContent(widget.NewVBox(input, button))
	w.ShowAndRun()
}
