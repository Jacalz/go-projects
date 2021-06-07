package main

import (
	"regexp"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("SEABorne")

	input := widget.NewEntry()
	input.SetPlaceHolder("Enter regex pattern...")

	button := &widget.Button{Text: "Find matches", Icon: theme.ContentCopyIcon()}
	button.OnTapped = func() {
		pattern, err := regexp.Compile(input.Text)
		if err != nil {
			dialog.ShowError(err, w)
			return
		}

		clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()

		result := pattern.FindAllString(clipboard.Content(), -1)
		items := len(result)
		if items == 0 {
			dialog.ShowInformation("No results", "Couldn't find anything with that expression.", w)
			return
		}

		var formated strings.Builder
		formated.Grow(items)
		for i := 0; i < items-1; i++ {
			formated.WriteString(result[i])
			formated.WriteString("\n")
		}

		formated.WriteString(result[items-1])
		clipboard.SetContent(formated.String())

		go func() {
			button.SetIcon(theme.ConfirmIcon())
			time.Sleep(200 * time.Millisecond)
			button.SetIcon(theme.ContentCopyIcon())
		}()
	}

	w.Resize(fyne.NewSize(400, 200))
	w.SetContent(container.NewVBox(input, button))
	w.ShowAndRun()
}
