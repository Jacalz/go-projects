package main

import (
	"fmt"
	"math"
)

// Konstanter för användning i beräkningarna.
const (
	Load         = 2000    //kg
	Length       = 1000000 //mm
	Gravity      = 9.82    //m/s²
	SafetyFactor = 2
)

// Beräkna arean för tråden från matematiska formeln.
func threadArea(ThreadDiameter float64) float64 {
	return (math.Pi * math.Pow(ThreadDiameter, 2)) / 4
}

// Beräkna trådens vikt utifrån matematiska formeln för vikt.
func threadMass(ThreadDiameter, Dencity float64) float64 {
	return (threadArea(ThreadDiameter) / 1000000) * Dencity * (Length / 1000)
}

// Beräkning av trådens tillåtna spänning.
func allowedThreadStrenght(AllowedThreadStrength, ThreadArea, ThreadMass float64) float64 {
	return ((AllowedThreadStrength * ThreadArea) / Gravity) - ThreadMass
}

// Beräkning av trådens sträckgräns.
func threadYieldStength(ThreadYieldStength, ThreadArea, ThreadMass float64) float64 {
	return ((ThreadYieldStength * ThreadArea) / Gravity) - ThreadMass
}

// Beräkning av trådens brottgräns.
func threadTensileStrength(ThreadTensileStrength, ThreadArea, ThreadMass float64) float64 {
	return ((ThreadTensileStrength * ThreadArea) / Gravity) - ThreadMass
}

// Antal trådar som krävs för att lyfta vikten.
func requiredThreads(ThreadStrengthValue float64) float64 {
	return math.Ceil(Load / ThreadStrengthValue)
}

// Beräkning av diameter på vajern utifrån arean.
func wireDiameter(WireArea float64) float64 {
	return math.Sqrt((WireArea * 4) / math.Pi)
}

// Totala förlängningen vid belastning.
func totalExtension(WireMass, WireArea, Elasticity float64) float64 {

	// Tyngden för vajern och belastningen.
	Q := WireMass * Gravity
	LoadF := Load * Gravity

	// Förlängning för vajer på grund av vajerns vik respective belastningens vikt.
	δQ := (Q * Length) / (2 * WireArea * Elasticity)
	δF := (LoadF * Length) / (WireArea * Elasticity)

	return δF + δQ
}

// Längden på trumman när man rullar upp vajern i ett lager.
func spoolLength(SpoolDiameter, WireDiameter, TotalExtension float64) float64 {
	LengthExtended := Length + TotalExtension

	// Diametern på trumman plus diametern på vajern för att få mitten av vajern på varje sida.
	TotalDiameter := SpoolDiameter + WireDiameter

	// Omkretsen för hur mycket vajer som går på ett varv följt av antalet varv som krävs för att rulla upp vajern.
	Circumreference := TotalDiameter * math.Pi
	Laps := LengthExtended / Circumreference

	return Laps * WireDiameter
}

// Längden på trumman när man rullar upp vajern i tre lager.
func threeLapsSpoolLength(SpoolDiameter, WireDiameter, TotalExtension float64) float64 {
	LengthExtended := Length + TotalExtension

	// Totala diameternara för varje varv upp till tre, varje varv ska vara två * radien av vajern större för add få mittpunkten av vajer då den är lika stor vid böjning.
	TotalDiameterLap := SpoolDiameter + WireDiameter
	TotalDiameter2Laps := SpoolDiameter + (3 * WireDiameter)
	TotalDiameter3Laps := SpoolDiameter + (5 * WireDiameter)

	Circumreference := ((math.Pi * TotalDiameterLap) + (math.Pi * TotalDiameter2Laps) + (math.Pi * TotalDiameter3Laps))

	Laps := LengthExtended / Circumreference

	return Laps * WireDiameter
}

// Material innehåller fält för alla egenskaper för materialet.
type Material struct {
	Elasticity      float64 //N/mm²
	Dencity         float64 //kg/m3
	YieldStrength   float64 //N/mm²
	TensileStrength float64 //N/mm²
	ThreadDiameter  float64 //mm
	SpoolDiameter   float64 //mm
}

func main() {

	// Alla egenskaper för materialen utifrån Material-typen.
	titan := Material{105000, 4600, 747.5, 962.5, 2.5, 500}
	cfrp := Material{107000, 1550, 800, 800, 2.5, 500}
	nylon := Material{2910, 1130, 72.4, 127.5, 5, 2000}

	var input string
	fmt.Print("\nMöjliga material är titan, cfrp eller nylon.\nVälj material att använda: ")
	fmt.Scanln(&input)

	// Skapa ett val av material där materialets egenskaper kopieras till material.
	var material *Material
	switch input {
	case "titan":
		material = &titan
	case "cfrp":
		material = &cfrp
	case "nylon":
		material = &nylon
	default:
		fmt.Printf("Error: %s är inte ett giltigt material!\n", input)
		return
	}

	// Beräkning av tillåten spänning samt definiering av trådens area och vikt.
	σTill := material.YieldStrength / SafetyFactor
	ThreadArea := threadArea(material.ThreadDiameter)
	ThreadMass := threadMass(material.ThreadDiameter, material.Dencity)

	// Skriv ut säkerhetsfaktorn och tråddiametern.
	fmt.Println("\nSäkerhetsfaktor:", SafetyFactor)
	fmt.Println("Diameter på lös tråd:", material.ThreadDiameter, "mm")

	// Definiering av variabler för tillåten styrka, sträckgräns och brottgräns utifrån funktionerna.
	AllowedThreadStrength := allowedThreadStrenght(σTill, ThreadArea, ThreadMass)
	ThreadYieldStength := threadYieldStength(material.YieldStrength, ThreadArea, ThreadMass)
	ThreadTensileStrength := threadTensileStrength(material.TensileStrength, ThreadArea, ThreadMass)

	// Definiering av variable för antalet trådar utifrån funktionen.
	RequiredThreads := requiredThreads(AllowedThreadStrength)

	// Skriv ut belastningsvikten och antalet trådar som krävs
	fmt.Println("\nBelastningsvikt:", Load)
	fmt.Println("Trådar som krävs:", RequiredThreads, "st")

	// Konvertera vikt, area och diameter från tråden över till hela vajern med hjälp av antalet trådar som krävs och funktionerna ovan.
	WireMass := RequiredThreads * ThreadMass
	WireArea := ThreadArea * RequiredThreads
	WireDiameter := wireDiameter(WireArea)

	// Skriv ut hela vajerns vikt.
	fmt.Printf("Vajerns vikt: %.2f kg\n", WireMass)

	// Konvertera trådens tillåtna styrka, sträckgräns och brottgräns till vajer.
	WireStrengthSafe := RequiredThreads * AllowedThreadStrength
	WireStrengthYield := RequiredThreads * ThreadYieldStength
	WireStrengthTensile := RequiredThreads * ThreadTensileStrength

	// Skriv ut värdena för tillåtna sträckgränsen, sträckgränsen och brottgräns.
	fmt.Printf("\nSäker sträckgräns: %.2f kg\n", WireStrengthSafe)
	fmt.Printf("Teoretisk sträckgräns: %.2f kg\n", WireStrengthYield)
	fmt.Printf("Teoretisk brottgräns: %.2f kg\n", WireStrengthTensile)

	// Definiera totala förlängningen och skriv ut den.
	TotalExtension := totalExtension(WireMass, WireArea, material.Elasticity)
	fmt.Printf("\nTotal förlängning: %.2f mm\n", TotalExtension)

	// Beräkna procentuella förlängningen.
	Epsilon := 100 * (TotalExtension / Length)
	fmt.Printf("Procentuell förlängning: %.5f%%\n", Epsilon)

	// Definiering av längen av trumman vid ett varv och tre varv samt utskrift av svaren, inklusive trummans diameter.
	SpoolLength := spoolLength(material.SpoolDiameter, WireDiameter, TotalExtension)
	ThreeLapsSpoolLength := threeLapsSpoolLength(material.SpoolDiameter, WireDiameter, TotalExtension)
	fmt.Printf("\nDiameter på trumman: %.2f mm\n", material.SpoolDiameter)
	fmt.Printf("Längd på trumman (ett lager): %.2f mm\n", SpoolLength)
	fmt.Printf("Längd på trumman (tre lager): %.2f mm\n", ThreeLapsSpoolLength)

}
