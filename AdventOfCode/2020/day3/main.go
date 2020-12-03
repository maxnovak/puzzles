package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("This is day 3")
	parsedMap := readFile()
	fmt.Println(parsedMap)
	result := findTreeCollisions(parsedMap)
	fmt.Println(result)
}

func findTreeCollisions(geologyMap [][]bool) int {
	x := 0
	collisions := 0
	for i, row := range geologyMap {
		if i == 0 {
			continue
		}
		x = (x + 3) % len(row)
		if row[x] {
			collisions++
		}
	}
	return collisions
}

func readFile() [][]bool {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var geologyMap [][]bool
	for scanner.Scan() {
		var xCoords []bool
		for _, letter := range scanner.Text() {
			xCoords = append(xCoords, letter == '#')
		}
		geologyMap = append(geologyMap, xCoords)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return geologyMap
}
