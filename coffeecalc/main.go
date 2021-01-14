package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

type dosage struct {
	water          float64
	measuringSpoon float64
	tableSpoon     float64
	teaSpoon       float64
}

func showHelp() {
	fmt.Println("Usage: coffeecalc cups")
	fmt.Println("\tcups must be a float between 1 and 10 inclusive with increments of 0.5")
	os.Exit(1)
}

func main() {

	var coffeetable = map[float64]dosage{
		1:   dosage{0.15, 1, 1, 2.5},
		1.5: dosage{0.2, 1.5, 1.5, 3.5},
		2:   dosage{0.25, 1.5, 1.5, 4.5},
		2.5: dosage{0.35, 2, 2, 6},
		3:   dosage{0.4, 2.5, 2.5, 7},
		3.5: dosage{0.45, 3, 3, 8},
		4:   dosage{0.5, 3, 3, 9},
		4.5: dosage{0.6, 3.5, 3.5, 10.5},
		5:   dosage{0.65, 4, 4, 11.5},
		5.5: dosage{0.7, 4.5, 4.5, 12.5},
		6:   dosage{0.75, 4.5, 4.5, 13.5},
		6.5: dosage{0.85, 5, 5, 15},
		7:   dosage{0.9, 5.5, 5.5, 16},
		7.5: dosage{0.95, 6, 6, 17},
		8:   dosage{1, 6, 6, 18},
		8.5: dosage{1.1, 6.5, 6.5, 19.5},
		9:   dosage{1.15, 7, 7, 20.5},
		9.5: dosage{1.2, 7.5, 7.5, 21.5},
		10:  dosage{1.25, 7.5, 7.5, 22.5},
	}

	if len(os.Args) < 2 {
		showHelp()
	}

	cups, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalf("Could not parse argument of %s to float", os.Args[1])
	}

	if cups < 1 || cups > 10 {
		showHelp()
	}

	this, exists := coffeetable[cups]
	if !exists {
		showHelp()
	}

	fmt.Printf("You need %.2f litres of water, and %.1f tableSpoons of coffee.", this.water, this.tableSpoon)
}
