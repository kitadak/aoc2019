package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
)

func CalculateFuel(module int) int {
	return int(math.Floor(float64(module)/3)) - 2
}

func CalculateFuelForFuel(fuel int) int {
	totalFuel := 0
	for fuel > 0 {
		fuel = CalculateFuel(fuel)
		if fuel < 0 {
			break
		}
		totalFuel += fuel
	}
	return totalFuel
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Usage: ./main <input_file>, got %d", len(os.Args))
	}

	dataset_file := os.Args[1]
	fmt.Println(dataset_file)
	file, err := os.Open(dataset_file)
	if err != nil {
		log.Fatal("Failed to load file", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	totalFuelForModules := 0
	// Part 1
	totalFuelForFuel := 0
	for scanner.Scan() {
		var module int
		fmt.Sscanf(scanner.Text(), "%d", &module)
		fuelForModule := CalculateFuel(module)
		totalFuelForModules += fuelForModule
		totalFuelForFuel += CalculateFuelForFuel(fuelForModule)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Total fuel for modules (part 1): %d\n", totalFuelForModules)
	fmt.Printf("Total fuel for fuel: %d\n", totalFuelForFuel)
	fmt.Printf("Total fuel for modules and fuel (part 2): %d\n", totalFuelForModules+totalFuelForFuel)
}
