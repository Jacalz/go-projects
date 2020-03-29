package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

// Create creates a file to download data to. It will save to specified file path if output is true, otherwice it will save to pwd.
func create(outpath string, output bool, url string) (out *os.File, err error) {

	// Save to current working directory if output is false.
	if !output {

		// Get the present working directory.
		pwd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		// Parse filename from inputed url and save to a variable.
		filename := filepath.Base(url)

		// Join the present working directory with the filename.
		outpath = pwd + "/" + filename
	}

	// Create the file where we should save our downloaded data.
	out, err = os.Create(outpath)
	if err != nil {
		return nil, err
	}

	return out, nil
}

// Download downloads data from an online source in to created file from create function.
func download(url string, outpath string, output bool) (err error) {
	// Run function to create the file we should save to.
	out, err := create(outpath, output, url)
	if err != nil {
		return err
	}

	// Make sure that we close the file when we are done. Should save us some memory.
	defer out.Close()

	// Download actual data from the website.
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	// Close the download body to save some memory.
	defer resp.Body.Close()

	// Write downloaded body to the file we created on the file system.
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func main() {

	// Parse all the flags and arguments.
	flag.Parse()

	// Set up the string avriables for url to fetch, flag for output and the resulting output path.
	var url string = flag.Arg(0)
	var outflag string = flag.Arg(1)
	var path string = flag.Arg(2)

	// Handle a bool for the outflag.
	var output bool
	if outflag == "--output" || outflag == "-output" {
		output = true
	} else if outflag == "-o" || outflag == "-O" {
		output = true
	} else {
		output = false
	}

	// Handle exceptions.
	if url == "" || url == "help" {
		log.Fatalln("Usage: get [url to file]\nOr:    get [url to file] -o [path to save to]")
	} else if output == true && path == "" {
		log.Fatalln("In order to specify output path, you need to add a path after -o:\nUsage: get [url to file] -o/-O/--output/-output [path to save to]")
	}

	// Download specified url and stop if we get an error.
	err := download(url, path, output)
	if err != nil {
		log.Fatal(err)
	}
}
