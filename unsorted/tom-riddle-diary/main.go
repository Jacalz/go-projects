package main

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// writeToDiary writes data to the diray and creats the file if it does not exist.
func writeToDiary(output string) {
	// Create the diary file on the file system.
	file, err := os.Create("diary.txt")
	if err != nil {
		log.Fatal("Could not create file")
	}

	// Write to the file.
	_, err = io.WriteString(file, output)
	if err != nil {
		log.Fatal("Failed to write to file")
	}
}

// readFromDiary reads the input of what is written inside the diary.
func readFromDiary() string {
	// Read the data from the diary.
	data, err := ioutil.ReadFile("diary.txt")
	if err != nil {
		log.Fatal("File reading error", err)
	}

	return string(data)
}

func main() {

	// Create the array with the output sentences.
	output := [4]string{
		"Hello and welcome to my diary!\nDo you want to hear a story?",
		"Wonderful!\nOnce up on a time I was a student at Hogwarts, and I wrote down my thoughts in this diary.\nNow I pass it along to you. Are you ready?",
		"Well then, let us begin!\nMy name is Tom Riddle and you have found my diary.\nYou will help me get revenge! What do you think of that?",
		"I'm afraid that you don't have a choice, my dear.\nYou will help me kill Harry Potter!",
	}

	// Create the array with the input sentences.
	input := [4]string{
		"Yes",
		"Yes, of course I am",
		"I don't wanna help you!",
		"I will obey",
	}

	// The loop that serves as the main point in our game.
	for i := 0; i < 4; i++ {

		// Write the output to the diary file. On first iteration, this will create the file.
		writeToDiary(output[i])

		// Sleep the main thread for five seconds.
		time.Sleep(5 * time.Second)

		// Clear the diary by printing "nothing" to it.
		writeToDiary("")

		// Sleep the main thread for 15 seconds while waiting for input.
		time.Sleep(15 * time.Second)

		// Read the text from the diary.
		text := readFromDiary()

		// Check that inputed text matches input data in array.
		if input[i] == text || input[i]+"\n" == text {
			continue
		} else {
			writeToDiary("You have entered an invalid answer! Please try again :)")
			break
		}
	}
}
