package main

import (
	"net/url"
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

	var formated strings.Builder
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

		formated.Grow(items)
		for i := 0; i < items-1; i++ {
			formated.WriteString(result[i])
			formated.WriteString("\n")
		}

		formated.WriteString(result[items-1])
		clipboard.SetContent(formated.String())
		formated.Reset()

		go func() {
			button.SetIcon(theme.ConfirmIcon())
			time.Sleep(200 * time.Millisecond)
			button.SetIcon(theme.ContentCopyIcon())
		}()
	}

	link, _ := url.Parse("https://cheatography.com/davechild/cheat-sheets/regular-expressions/")
	cheat := &widget.Hyperlink{Text: "Regex Cheat Sheet", URL: link}

	w.Resize(fyne.NewSize(400, 200))
	w.SetContent(container.NewVBox(input, button, cheat))
	w.ShowAndRun()
}
