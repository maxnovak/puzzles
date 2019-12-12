package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Planet struct {
	Name   string
	Orbits *Planet
}

func main() {
	orbits := readFile()

	for _, planet := range orbits {
		fmt.Println(planet.Name, planet.Orbits.Name)
	}

}

func readFile() []Planet {

	file, err := os.Open("day-6-input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var orbits []Planet

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		planets := strings.Split(scanner.Text(), ")")
		center := planets[0]
		outer := Planet{Name: planets[1]}

		orbits = append(orbits, Planet{Name: center, Orbits: &outer})
	}

	return orbits
}
