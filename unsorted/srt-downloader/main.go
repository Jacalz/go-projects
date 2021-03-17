package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"strings"
)

// https:\/\/prod.cdn.bbaws.net\/TDC_Blockbuster_-_Production\/995\/64\/WB_6000124948_swe.srt
const srtURLExp = `https:\\/\\/prod\.cdn\.bbaws\.net\\/TDC_Blockbuster_-_Production\\/\d{1,20}\\/\d{1,20}\\/\w{1,20}_\d{1,20}_swe\.srt`

var srtMatcher = regexp.MustCompile(srtURLExp)

func getSrtURL(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	scan := bufio.NewScanner(resp.Body)
	scan.Split(bufio.ScanLines)

	rawSrtURL := ""
	for scan.Scan() {
		raw := srtMatcher.FindString(scan.Text())
		if raw != "" {
			rawSrtURL = raw
			break
		}
	}

	if rawSrtURL == "" {
		return "", errors.New("no match could be found")
	}

	return strings.ReplaceAll(rawSrtURL, `\`, ""), nil
}

func saveSRTFile(url, filename string) error {
	srt, err := http.Get(url)
	if err != nil {
		return err
	}
	defer srt.Body.Close()

	srtFile, err := os.Create(filename)
	if err != nil {
		return err
	}

	if _, err = io.Copy(srtFile, srt.Body); err != nil {
		return err
	}

	fmt.Println("Saved the SRT file as:", filename)
	return nil
}

func main() {
	url := ""
	fmt.Print("Enter url to grab: ")
	fmt.Scanln(&url)

	srtURL, err := getSrtURL(url)
	if err != nil {
		log.Fatalln("Error on finding SRT file URL:", err)
	}

	if err = saveSRTFile(srtURL, path.Base(url)+".srt"); err != nil {
		log.Fatalln("Error on saving the SRT file:", err)
	}
}
