package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Print("Press enter to start!")
	fmt.Scanln()

	// Sleep the main thread for a random number of seconds between 1 and 15
	time.Sleep(time.Duration(rand.Int63n(15)) * time.Second)

	// Take the time now and print a box
	start := time.Now()
	row := "<------------->"
	fmt.Printf("			%s\n			%s\n			%s\n			%s\n", row, row, row, row)

	// When user presses enter, a new time is taken and then we calculate the time between those points.
	fmt.Scanln()
	end := time.Now()
	elapsed := end.Sub(start)

	fmt.Print("Your reaction time is: ", elapsed)

}
