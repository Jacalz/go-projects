package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

// The fetch command runs the git clone command and prints the output to terminal.
func fetch(command string) {
	fetch := exec.Command("sh", "-c", command)
	output, err := fetch.CombinedOutput()
	if err != nil {
		log.Fatalf("%s", output)
	} else {
		fmt.Printf("%s", output)
	}
}

func main() {
	// Let flag parse the command line arguments and then input the first argument in to the repo variable.
	flag.Parse()
	repo := flag.Arg(0)

	// Handle empty arguments by printing usage for out program.
	if repo == "" {
		log.Fatalln("Usage: solfetch [repository name]")
	}

	// Run the git command to download the whole repository from upstream.
	fetch(fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", repo))
}
