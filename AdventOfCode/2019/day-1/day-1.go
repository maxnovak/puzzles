package main

import (
	"fmt"
	"math"
)

func main() {
	//Test cases
	fmt.Println("Mass of 12: ", calculateMass(12))
	fmt.Println("Mass of 14: ", calculateMass(14))
	fmt.Println("Mass of 1969: ", calculateMass(1969))
	fmt.Println("Mass of 100756: ", calculateMass(100756))

}

func calculateMass(mass float64) float64 {
	fuel := math.Floor(mass / 3) - 2
	return fuel
}