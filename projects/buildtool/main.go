package main

import (
	"flag"
	"fmt"
	"log"
	"os/exec"
	"path"
)

// Use PGO by compiling with -fprofile-generate, running and than recompiling with -fprofile-use and -profile-correction
const flags = "-O3 -march=native -mtune=native -flto -ffast-math -pipe -Wall -v"

func match(pattern, input string) bool {
	output, _ := path.Match(pattern, input)
	return output
}

func runCompile(command string) {
	compile := exec.Command("sh", "-c", command)
	output, err := compile.CombinedOutput()
	if err != nil {
		log.Fatalf("%s\n", output)
	} else {
		fmt.Printf("%s\n", output)
	}
}

func compileCPP(file string) {
	command := "clang++ -std=c++14 -stdlib=libc++ " + file + " -o compiled-cpp " + flags
	runCompile(command)
}

func compileGo(file string) {
	command := `go build -ldflags="-s -w" -o compiled-go ` + file
	runCompile(command)
}

func checkLang(input string) {
	if match("*.cpp", input) || match("*.cxx", input) {
		compileCPP(input)
	} else if match("*.go", input) {
		compileGo(input)
	} else {
		fmt.Println("Usage:\n   build [file-to-build]\n",
			"     Valid languages:\n",
			"       C++ - Uses same optimizing compiler flags as C in Clang++ compiler.\n",
			"       Go  - Builds without debug and dwarf data.\n",
			"\n       File to build:\n",
			"       Specify file - Filetypes .go, .cpp and .cxx are supported.",
		)
	}
}

func main() {
	flag.Parse()
	input := flag.Arg(0)
	checkLang(input)
}
