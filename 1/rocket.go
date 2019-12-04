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
	totalFuel := 0
	for scanner.Scan() {
		var module int
		fmt.Sscanf(scanner.Text(), "%d", &module)
		totalFuel += CalculateFuel(module)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(totalFuel)
}
