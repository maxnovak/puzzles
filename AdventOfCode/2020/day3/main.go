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
	result11 := findTreeCollisions(parsedMap, 1, 1)
	result31 := findTreeCollisions(parsedMap, 3, 1)
	result51 := findTreeCollisions(parsedMap, 5, 1)
	result71 := findTreeCollisions(parsedMap, 7, 1)
	result12 := findTreeCollisions(parsedMap, 1, 2)
	fmt.Printf("Right 1, down 1: %v\n", result11)
	fmt.Printf("Right 3, down 1: %v\n", result31)
	fmt.Printf("Right 5, down 1: %v\n", result51)
	fmt.Printf("Right 7, down 1: %v\n", result71)
	fmt.Printf("Right 1, down 2: %v\n", result12)

	fmt.Printf("Answer to part 2: %v", result11*result12*result31*result51*result71)
}

func findTreeCollisions(geologyMap [][]bool, rightAmount int, downAmount int) int {
	x := 0
	collisions := 0
	lengthOfData := len(geologyMap[0])
	for y := downAmount; y < len(geologyMap); y = y + downAmount {
		x = (x + rightAmount) % lengthOfData
		if geologyMap[y][x] {
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
