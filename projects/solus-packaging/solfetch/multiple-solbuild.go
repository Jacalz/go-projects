package main

import (
	//"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

// exists returns whether the given file or directory exists
func exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// The fetch command runs the git clone command and prints the output to terminal.
func fetch(command string) []byte {
	fetch := exec.Command("sh", "-c", command)
	output, err := fetch.CombinedOutput()
	if err != nil && command != "" {
		panic(err)
	}

	return output
}

func main() {
	// Parse three first arguments in to repo slice.
	//flag.Parse()
	//repo := [3]string{flag.Arg(0), flag.Arg(1), flag.Arg(2)}
	repo := [3]string{"atom", "brave", "mutter"}

	// Don't proceed if they alread are fetched.
	if exists(repo[0]) || exists(repo[1]) || exists(repo[2]) {
		log.Fatalln("Don't fetch repos that are already fetched!")
	}

	// Handle empty arguments and tell user how to use program.
	if repo[0] == "" {
		log.Fatalln("Usage: solfetch [repository name] [optional] [optional]")
	}

	var one, two, three []byte

	if repo[0] != "" {
		one = fetch(fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", repo[0]))
	}
	if repo[1] != "" {
		two = fetch(fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", repo[1]))
	}
	if repo[2] != "" {
		three = fetch(fmt.Sprintf("git clone https://dev.getsol.us/source/%s.git", repo[2]))
	}

	// Print output to terminal:
	fmt.Printf("%s%s%s", one, two, three)
}
