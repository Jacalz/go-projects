package main

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// wordBoundry is an extension for the regex word boundry.
// Combined with, clean(), it allows us to match many more serial numbers.
const wordBoundry = `(\b|\n)`

// Clean trims any leading newlines (trailing are handled later).
func clean(input string) string {
	return strings.TrimPrefix(input, "\n")
}

func containsDuplicates(input []string) bool {
	dup := make(map[string]bool)

	for _, item := range input {
		if _, ok := dup[item]; ok {
			return true
		} else {
			dup[item] = true
		}
	}

	return false
}

func main() {
	a := app.New()
	a.SetIcon(resourceIconPng)
	w := a.NewWindow("SEABorne")

	input := &widget.Entry{PlaceHolder: "Enter regex pattern..."}

	word := &widget.Check{Text: "Match whole words only", Checked: true}

	const defaultInfo = "Objects found: 0\nDuplicated items: no"
	info := &widget.Label{Text: defaultInfo}

	var formated strings.Builder
	button := &widget.Button{Text: "Find matches", Icon: theme.ContentCopyIcon()}
	button.OnTapped = func() {
		pattern := input.Text
		if word.Checked {
			pattern = wordBoundry + pattern + wordBoundry
		}

		expression, err := regexp.Compile(pattern)
		if err != nil {
			info.SetText(defaultInfo)
			dialog.ShowError(err, w)
			return
		}

		clipboard := fyne.CurrentApp().Driver().AllWindows()[0].Clipboard()

		result := expression.FindAllString(clipboard.Content(), -1)
		items := len(result)
		if items == 0 {
			info.SetText(defaultInfo)
			dialog.ShowInformation("No matches found", "Could not find anything with that expression.", w)
			return
		}

		formated.Grow(items)
		for i := 0; i < items-1; i++ {
			cleaned := clean(result[i])
			formated.WriteString(cleaned)

			if !strings.HasSuffix(cleaned, "\n") {
				formated.WriteString("\n")
			}
		}

		info.Text = "Objects found: " + strconv.Itoa(items) + "\nDuplicated items: no"
		if containsDuplicates(result) {
			info.SetText(info.Text[:len(info.Text)-2] + "yes")
		}

		formated.WriteString(clean(result[items-1]))
		clipboard.SetContent(formated.String())
		formated.Reset()

		go func() {
			button.SetIcon(theme.ConfirmIcon())
			time.Sleep(200 * time.Millisecond)
			button.SetIcon(theme.ContentCopyIcon())
		}()
	}

	input.OnSubmitted = func(_ string) {
		button.OnTapped()
	}

	link, _ := url.Parse("https://cheatography.com/davechild/cheat-sheets/regular-expressions/")
	cheat := &widget.Hyperlink{Text: "Regex Cheat Sheet", URL: link}

	w.Resize(fyne.NewSize(450, 250))
	w.SetContent(container.NewVBox(input, button, word, info, cheat))
	w.ShowAndRun()
}
