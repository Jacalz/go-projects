package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
)

// Exists returns whether the given file or directory exists
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}

	return false
}

// The fetch command runs the git clone command and prints the output to terminal.
func fetch(command string, finished chan struct{}) {
	fetch := exec.Command("sh", "-c", fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", command))
	output, err := fetch.CombinedOutput()
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%s", output)
	}

	close(finished)
}

func main() {
	// Grab an array of all our arguments.
	flag.Parse()
	repo := flag.Args()

	// Don't proceed if any any of the repos are already fetched.
	for _, name := range repo {
		if exists(name) {
			fmt.Println("Please check that you are not fetching repos that are already fetched.")
			return
		}
	}

	// Handle empty arguments and tell user how to use program.
	if repo[0] == "" {
		fmt.Println("Usage: solfetch [repository name] [any amount of optional repos]...")
		return
	}

	// Spin up all of our communication channels.
	var channel []chan struct{}
	for range repo {
		channel = append(channel, make(chan struct{}))
	}

	// Start up cocurrent tasks for fetching all repos at once.
	for i := range repo {
		go fetch(repo[i], channel[i])
	}

	// Loop through and grab the output of each boolean channel without storing it.
	for i := range repo {
		<-channel[i]
	}
}
