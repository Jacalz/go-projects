package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"os"
)

type Materials struct {
	XMLName   xml.Name   `xml:"materials"`
	Materials []Material `xml:"material"`
}

type Material struct {
	XMLName    xml.Name   `xml:"material"`
	Name       string     `xml:"name"`
	Properties Properties `xml:"properties"`
}

type Properties struct {
	XMLName         xml.Name `xml:"properties"`
	Elasticity      float64  `xml:"elasticity"`
	Dencity         float64  `xml:"dencity"`
	Yieldstrength   float64  `xml:"yieldstrength"`
	Tensilestrength float64  `xml:"tensilestrength"`
	Load            float64  `xml:"load"`
	Threaddiameter  float64  `xml:"threaddiameter"`
	Spooldiameter   float64  `xml:"spooldiameter"`
	Wirelength      float64  `xml:"wirelength"`
	Safetyfactor    float64  `xml:"safetyfactor"`
}

func parse() Materials {
	file, err := os.Open("material.xml")
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}

	byteValue, _ := ioutil.ReadAll(file)
	var materials Materials
	xml.Unmarshal(byteValue, &materials)

	return materials
}
