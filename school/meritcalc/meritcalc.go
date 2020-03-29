package main

import (
	"fmt"
	"log"
)

func betyg(bokstav string) (nummer float64) {

	switch bokstav {
	case "A":
		nummer = 20
	case "B":
		nummer = 17.5
	case "C":
		nummer = 15
	case "D":
		nummer = 12.5
	case "E":
		nummer = 10
	case "F":
		nummer = 0
	default:
		log.Fatalln("Inte ett giltigt betyg!")
	}

	return nummer
}

func merit(antal int, extra float64) (merit float64) {
	fmt.Println("Skriv in ett betyg per rad:")
	var bokstav string

	for i := 1; i <= antal; i++ {
		fmt.Scanln(&bokstav)
		merit += betyg(bokstav)
	}

	merit = (merit / float64(antal)) + extra

	return merit
}

func main() {
	var betyg int
	var extra float64
	fmt.Print("Ange antal betyg från gymnasiet: ")
	fmt.Scanln(&betyg)

	fmt.Print("Antal extramerit från kurser: ")
	fmt.Scanln(&extra)

	if extra < 0 || extra > 2.5 {
		log.Fatalln("Inte ett giltigt antal extramerit!")
	}

	fmt.Printf("Meritpoäng: %.2f", merit(betyg, extra))
}
