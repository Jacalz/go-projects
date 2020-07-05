package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sync"
)

func fetch(repo string, wg *sync.WaitGroup) {
	cmd := exec.Command("git", "clone", "https://dev.getsol.us/source/"+repo)
	cmd.Stdout, cmd.Stderr = os.Stdout, os.Stderr

	cmd.Run()
	wg.Done()
}

func main() {
	flag.Parse()
	repos := flag.Args()

	if repos[0] == "" {
		fmt.Println("Usage: solfetch [repository name] [any amount of optional repos]...")
		return
	}

	var wg sync.WaitGroup

	for _, v := range repos {
		go fetch(v, &wg)
		wg.Add(1)
	}

	wg.Wait()
}
