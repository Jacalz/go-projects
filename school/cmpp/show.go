package main

import (
	"fmt"
	"log"
)

// PrintLessThan print all names with less than x visits.
func PrintLessThan(lessthan int) {
	data := ReadDataFromXML()
	for i := 0; i < len(data.Person); i++ {
		if data.Person[i].Visits < lessthan {
			fmt.Println(data.Person[i].Name, data.Person[i].Visits)
		}
	}
}

// CheckNumber checks for inputed number from command less.
func CheckNumber(command string) (lessthan int) {
	switch command {
	case "10":
		lessthan = 10
	case "9":
		lessthan = 9
	case "8":
		lessthan = 8
	case "7":
		lessthan = 7
	case "6":
		lessthan = 6
	case "5":
		lessthan = 5
	case "4":
		lessthan = 4
	case "3":
		lessthan = 3
	case "2":
		lessthan = 2
	default:
		log.Fatalln("Enter a value from two up to ten.")
	}

	return
}
