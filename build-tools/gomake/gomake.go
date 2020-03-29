package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

// compileRelease64 compiles a binary without debug symbols for a smaller size.
func compileRelease64() {
	fmt.Println("Compiling binaries without debug symbols...")
	cmd := exec.Command("/bin/bash", "-c", `go build -ldflags="-s -w"`)
	cmd.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

// crossCompileRelease64 compiles the binary for linux, macOS and windows and all without debug symbols.
func crossCompile64(Release bool) {

	// Build as debug if release isn't passed as flag and if it is we will build without debug symbols
	var LinuxCommand, DarwinCommand, WindowsCommand string
	if Release {
		fmt.Println("Compiling 64-bit binaries (no debug) for Linux, MacOS and Windows...")
		LinuxCommand = `go build -ldflags="-s -w" -o $(pwd)/main-linux-x64`
		DarwinCommand = `go build -ldflags="-s -w" -o $(pwd)/main-darwin-x64`
		WindowsCommand = `go build -ldflags="-s -w" -o $(pwd)/main-windows-x64`
	} else {
		fmt.Println("Compiling 64-bit binaries for Linux, MacOS and Windows...")
		LinuxCommand = "go build -o $(pwd)/main-linux-x64"
		DarwinCommand = "go build -o $(pwd)/main-darwin-x64"
		WindowsCommand = "go build -o $(pwd)/main-windows-x64"
	}

	// cmdOne has the command to build for linux, cmdOne.Env has the environment variables and then we run with cmdOne.Run.
	linux := exec.Command("/bin/bash", "-c", LinuxCommand)
	linux.Env = append(os.Environ(), "GOOS=linux", "GOARCH=amd64")
	if err := linux.Run(); err != nil {
		log.Fatal(err)
	}

	// cmdTwo has the command to build for macOS, cmdTwo.Env has the environment variables and then we run with cmdTwo.Run.
	darwin := exec.Command("/bin/bash", "-c", DarwinCommand)
	darwin.Env = append(os.Environ(), "GOOS=darwin", "GOARCH=amd64")
	if err := darwin.Run(); err != nil {
		log.Fatal(err)
	}

	// cmdThree has the command to build for windows, cmdThree.Env has the environment variables and then we run with cmdThree.Run.
	windows := exec.Command("/bin/bash", "-c", WindowsCommand)
	windows.Env = append(os.Environ(), "GOOS=windows", "GOARCH=amd64")
	if err := windows.Run(); err != nil {
		log.Fatal(err)
	}
}

// Just a standard compilation using go build.
func compile() {
	fmt.Println("Running standard compile...")
	compile := exec.Command("go", "build")
	if err := compile.Run(); err != nil {
		log.Fatal(err)
	}
}

func cleanupCrossCompiles() {
	fmt.Println("Cleaning up files from cross compiling...")
	remove := exec.Command("/bin/bash", "-c", "rm main-linux-x64 main-darwin-x64 main-windows-x64")
	if err := remove.Run(); err != nil {
		log.Fatal(err, "\nNo cross compiled packages to remove")
	}
}

func main() {

	// Check that user uses Linux to run the software.
	if runtime.GOOS != "linux" {
		log.Fatalf("%s is not a supported operatingsystem, use Linux instead!\n", runtime.GOOS)
	}

	// Set up command line flags with documentation and then parse them.
	var Release *bool = flag.Bool("release", false, "strip debug data for a smaller binary file.")
	var CrossCompile *bool = flag.Bool("cross-compile", false, "compiles the package for Amd64 on Linux, MacOS and Windows without debug symbols.")
	flag.Parse()
	var Clean string = flag.Arg(0)

	if Clean == "clean" {
		cleanupCrossCompiles()
	} else if *Release || *CrossCompile {

		// Handle values for the command line flags passed by users.
		if *Release && *CrossCompile == false {
			compileRelease64()
		} else if *CrossCompile {
			crossCompile64(*Release)
		} else {
			compile()
		}
	}
}
