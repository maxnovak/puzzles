package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	//Test cases
	fmt.Println("Mass of 12: ", calculateFuel(12))
	fmt.Println("Mass of 14: ", calculateFuel(14))
	fmt.Println("Mass of 1969: ", calculateFuel(1969))
	fmt.Println("Mass of 100756: ", calculateFuel(100756))

	//Sum Fuels from Masses
	masses := readFile()
	fuelSum := float64(0)
	for _, mass := range masses {
		fuelSum += calculateFuel(mass)
	}
	fmt.Printf("fuel sum: %f", fuelSum)
}

func calculateFuel(mass float64) float64 {
	fuel := math.Floor(mass / 3) - 2
	if (math.Floor(fuel / 3) - 2 > 0) {
		fuel += calculateFuel(fuel)
	}
	return fuel
}

func readFile() []float64{

	file, err := os.Open("day-1-input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var masses []float64
    for scanner.Scan() {
		number, _ := strconv.ParseFloat(scanner.Text(), 64)
        masses = append(masses, number)
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return masses
}