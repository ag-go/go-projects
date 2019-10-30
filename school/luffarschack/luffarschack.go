package main

import (
	"fmt"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

// Global variable for defining if a button is pressed or not.
var clicked = [9]int8{}

// Global variables for which parts the players own. Index 0 is the first tile and index 8 is the ninth tile.
var player1 = [9]bool{false, false, false, false, false, false, false, false, false}
var player2 = [9]bool{false, false, false, false, false, false, false, false, false}

// Our global index for defining if it's player one or player two's turn to play.
var index int8

// InitGUI starts up the whole interface for out program.
func InitGUI() {
	// Initialize our new fyne interface application.
	app := app.New()

	// Set the application icon for our program.
	app.SetIcon(icon)

	// Create the window for our user interface.
	window := app.NewWindow("Luffarschack")

	// Make a new vertical box where we can stack our horizontal grids.
	vbox := widget.NewVBox()

	// Define all variables we need for the buttons.
	var (
		button1 = widget.NewButton("", func() {
			clicked[0] = 1
		})
		button2 = widget.NewButton("", func() {
			clicked[1] = 1
		})
		button3 = widget.NewButton("", func() {
			clicked[2] = 1
		})
		button4 = widget.NewButton("", func() {
			clicked[3] = 1
		})
		button5 = widget.NewButton("", func() {
			clicked[4] = 1
		})
		button6 = widget.NewButton("", func() {
			clicked[5] = 1
		})
		button7 = widget.NewButton("", func() {
			clicked[6] = 1
		})
		button8 = widget.NewButton("", func() {
			clicked[7] = 1
		})
		button9 = widget.NewButton("", func() {
			clicked[8] = 1
		})
	)

	// Append a start button to our vertical box.
	vbox.Append(widget.NewButton("Click to start", func() {
		// Make sure to clear all earlier button presses before proceeding.
		for i := range clicked {
			clicked[i] = 0
		}

		// Also remove all our icons before proceeding if we are restarting.
		button1.SetIcon(nil)
		button2.SetIcon(nil)
		button3.SetIcon(nil)
		button4.SetIcon(nil)
		button5.SetIcon(nil)
		button6.SetIcon(nil)
		button7.SetIcon(nil)
		button8.SetIcon(nil)
		button9.SetIcon(nil)

		// Clear the index to make sure that we start from scratch.
		index = 0

		// Clear the board information for each player.
		player1 = [9]bool{false, false, false, false, false, false, false, false, false}
		player2 = [9]bool{false, false, false, false, false, false, false, false, false}

		// Define the variable for telling that we won.
		won := false

		for won == false {

			// Ugly switch statement to set the icons for our buttons.
			if clicked[0] == 1 {
				clicked[0]++
				if index%2 == 0 {
					button1.SetIcon(circle)
					player1[0] = true
				} else {
					button1.SetIcon(cross)
					player2[0] = true
				}
				index++
			} else if clicked[1] == 1 {
				clicked[1]++
				if index%2 == 0 {
					button2.SetIcon(circle)
					player1[1] = true
				} else {
					button2.SetIcon(cross)
					player2[1] = true
				}
				index++
			} else if clicked[2] == 1 {
				clicked[2]++
				if index%2 == 0 {
					button3.SetIcon(circle)
					player1[2] = true
				} else {
					button3.SetIcon(cross)
					player2[2] = true
				}
				index++
			} else if clicked[3] == 1 {
				clicked[3]++
				if index%2 == 0 {
					button4.SetIcon(circle)
					player1[3] = true
				} else {
					button4.SetIcon(cross)
					player2[3] = true
				}
				index++
			} else if clicked[4] == 1 {
				clicked[4]++
				if index%2 == 0 {
					button5.SetIcon(circle)
					player1[4] = true
				} else {
					button5.SetIcon(cross)
					player2[4] = true
				}
				index++
			} else if clicked[5] == 1 {
				clicked[5]++
				if index%2 == 0 {
					button6.SetIcon(circle)
					player1[5] = true
				} else {
					button6.SetIcon(cross)
					player2[5] = true
				}
				index++
			} else if clicked[6] == 1 {
				clicked[6]++
				if index%2 == 0 {
					button7.SetIcon(circle)
					player1[6] = true
				} else {
					button7.SetIcon(cross)
					player2[6] = true
				}
				index++
			} else if clicked[7] == 1 {
				clicked[7]++
				if index%2 == 0 {
					button8.SetIcon(circle)
					player1[7] = true
				} else {
					button8.SetIcon(cross)
					player2[7] = true
				}
				index++
			} else if clicked[8] == 1 {
				clicked[8]++
				if index%2 == 0 {
					button9.SetIcon(circle)
					player1[8] = true
				} else {
					button9.SetIcon(cross)
					player2[8] = true
				}
				index++
			}

			// Just print some debug info to the terminal, because we stall the gui otherwice.
			fmt.Println("Player 1:", player1, "Player 2:", player2)

			// Check if our index is bigger than 9, because then we are finished. If the index is bigger or equal to 4, we can check for a win.
			if index > 9 {
				break
			} else if index >= 4 {
				if CheckWon(player1) {
					// Show a dialogue informing the first player that he won!
					message := dialog.NewInformation("Player 1 has won!", "Congratulations to player 1 for winning.", window)
					message.Show()
					break
				} else if CheckWon(player2) {
					// Show a dialogue informing the second player that he won!
					message := dialog.NewInformation("Player 2 has won!", "Congratulations to player 2 for winning.", window)
					message.Show()
					break
				} else if index == 9 {
					message := dialog.NewInformation("It is a tie!", "Nobody has won. Please try better next time.", window)
					message.Show()
					break
				}
			}

		}
	}))

	// Append each new row as a new container with gird layout and three buttons.
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button1, button2, button3))
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button4, button5, button6))
	vbox.Append(fyne.NewContainerWithLayout(layout.NewGridLayout(3), button7, button8, button9))

	// Add our vertical box to be viewed.
	window.SetContent(vbox)

	// Set the size to something small but good looking and usable.
	window.Resize(fyne.NewSize(400, 100))

	// Show all of our set content and run the gui.
	window.ShowAndRun()
}

func main() {
	InitGUI()
}

// CheckWon checks all possible combinations for winning.
func CheckWon(player [9]bool) bool {

	// All possible combinations for winning.
	if player[0] && player[1] && player[2] {
		return true
	} else if player[0] && player[3] && player[6] {
		return true
	} else if player[0] && player[3] && player[5] {
		return true
	} else if player[6] && player[7] && player[8] {
		return true
	} else if player[2] && player[5] && player[8] {
		return true
	} else if player[0] && player[4] && player[8] {
		return true
	} else if player[2] && player[4] && player[6] {
		return true
	} else if player[1] && player[4] && player[7] {
		return true
	} else if player[3] && player[4] && player[5] {
		return true
	}

	return false
}