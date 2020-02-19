package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// Define all our global variables. We use these in order to not redeclare them every time we start the game.
var (
	// Variables for which parts the players own. Index 0 is the first tile and index 8 is the ninth tile.
	player1 [9]bool
	player2 [9]bool

	// Variable to handle if buttons are already pressed.
	pressed [9]bool

	// Tell if we are in game or not.
	inGame bool

	// Index for defining if it's player one or player two's turn to play. Use 8bit variable to save on memory allocation.
	index uint8
)

func main() {
	InitGUI()
}

// InitGUI starts up the whole interface for our program.
func InitGUI() {
	// Initialize our new fyne interface application.
	app := app.NewWithID("com.github.jacalz.tic-tac-toe")

	// Set the application icon for our program.
	app.SetIcon(icon)

	// Create the window for our user interface.
	window := app.NewWindow("Tic-Tac-Toe")

	// Set up a buffered channel for sending button presses.
	channel := make(chan uint8)

	// Create all our buttons tiles for the game and send a button number to our channel on each press.
	var (
		button1 = widget.NewButton("", func() {
			if !pressed[0] && inGame {
				channel <- 0
			}
		})
		button2 = widget.NewButton("", func() {
			if !pressed[1] && inGame {
				channel <- 1
			}
		})
		button3 = widget.NewButton("", func() {
			if !pressed[2] && inGame {
				channel <- 2
			}
		})
		button4 = widget.NewButton("", func() {
			if !pressed[3] && inGame {
				channel <- 3
			}
		})
		button5 = widget.NewButton("", func() {
			if !pressed[4] && inGame {
				channel <- 4
			}
		})
		button6 = widget.NewButton("", func() {
			if !pressed[5] && inGame {
				channel <- 5
			}
		})
		button7 = widget.NewButton("", func() {
			if !pressed[6] && inGame {
				channel <- 6
			}
		})
		button8 = widget.NewButton("", func() {
			if !pressed[7] && inGame {
				channel <- 7
			}
		})
		button9 = widget.NewButton("", func() {
			if !pressed[8] && inGame {
				channel <- 8
			}
		})
	)

	// Create our start button for the whole game.
	startButton := widget.NewButton("Click to start", func() {
		go func() {
			// Make sure that all buttons are reset to remove icons before we start.
			button1.SetIcon(nil)
			button2.SetIcon(nil)
			button3.SetIcon(nil)
			button4.SetIcon(nil)
			button5.SetIcon(nil)
			button6.SetIcon(nil)
			button7.SetIcon(nil)
			button8.SetIcon(nil)
			button9.SetIcon(nil)

			// Tell the rest of the program that we are in a game.
			inGame = true

			var message dialog.Dialog

			// The main loop for the game.
			for index = 0; index < 9; index++ {

				// Sleep the loop until we get a number in the channel.
				clicked := <-channel

				// Handles all our button presses during the play time.
				switch clicked {
				case 0:
					PressHandler(button1, 0)
				case 1:
					PressHandler(button2, 1)
				case 2:
					PressHandler(button3, 2)
				case 3:
					PressHandler(button4, 3)
				case 4:
					PressHandler(button5, 4)
				case 5:
					PressHandler(button6, 5)
				case 6:
					PressHandler(button7, 6)
				case 7:
					PressHandler(button8, 7)
				case 8:
					PressHandler(button9, 8)
				}

				// Check if index is bigger or equal to 4, because it's the earliest time we can win. If index is 8, we have a tie and nobody won.
				if index >= 4 {
					if CheckWon(player1) {
						// Show a dialogue informing the first player that he won!
						message = dialog.NewInformation("Player 1 has won!", "Congratulations to player 1 for winning.", window)
						break
					} else if CheckWon(player2) {
						// Show a dialogue informing the second player that he won!
						message = dialog.NewInformation("Player 2 has won!", "Congratulations to player 2 for winning.", window)
						break
					} else if index == 8 {
						// It is a tie if the game hasn't ended before index reaches 8 and no one wins on the ninth placement.
						message = dialog.NewInformation("It is a tie!", "Nobody has won. Please try better next time.", window)
					}
				}
			}

			// Clean up after our game finishes and do it on an other goroutine to speed it up.
			go func() {
				// We are not in a game anymore.
				inGame = false

				// As a clean up we make sure to clear all markers for each player and for all pressed buttons.
				player1 = [9]bool{false, false, false, false, false, false, false, false, false}
				player2 = [9]bool{false, false, false, false, false, false, false, false, false}
				pressed = [9]bool{false, false, false, false, false, false, false, false, false}
			}()

			// Show our winning or tie message to the user.
			message.Show()
		}()

	})

	// Add all the buttons in to a three column grid layout inside a container.
	container := fyne.NewContainerWithLayout(layout.NewGridLayout(3), layout.NewSpacer(), startButton, layout.NewSpacer(), button1, button2, button3, button4, button5, button6, button7, button8, button9)

	// Set the container as what is being displayed.
	window.SetContent(container)

	// Set a sane default for the window size.
	window.Resize(fyne.NewSize(400, 250))

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}

// PressHandler handles the press of a button and updates button icons and player arrays accordingly.
func PressHandler(button *widget.Button, num uint8) {
	// Check if player one or two presses the button and handle it accordingly.
	if index%2 == 0 {
		button.SetIcon(circle)
		player1[num] = true
	} else {
		button.SetIcon(cross)
		player2[num] = true
	}

	// Set the button as pressed to not make it pressable again.
	pressed[num] = true
}

// CheckWon checks all possible combinations for winning.
func CheckWon(player [9]bool) bool {

	// Switch statement with all possible combinations for winning.
	switch {
	case player[0] && player[1] && player[2]:
		return true
	case player[3] && player[4] && player[5]:
		return true
	case player[6] && player[7] && player[8]:
		return true
	case player[0] && player[3] && player[6]:
		return true
	case player[1] && player[4] && player[7]:
		return true
	case player[2] && player[5] && player[8]:
		return true
	case player[0] && player[4] && player[8]:
		return true
	case player[2] && player[4] && player[6]:
		return true
	}

	return false
}
