package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
)

func run(command string) {
	run := exec.Command("sh", "-c", fmt.Sprintf("sudo %s", command))
	output, err := run.CombinedOutput()
	if err != nil {
		log.Fatalf("%s", output)
	} else {
		fmt.Printf("%s", output)
	}
}

func main() {
	// Option to clean the local repo.
	var clean *bool = flag.Bool("clean", false, "Clear the local repo of eopkg files.")
	flag.Parse()

	if *clean {
		run("rm -fv /var/lib/solbuild/local/*.eopkg")
	} else {
		run("cp -v *.eopkg /var/lib/solbuild/local")
	}
}
