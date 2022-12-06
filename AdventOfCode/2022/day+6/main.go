package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	data := readFile()
	firstMarker := findFirstMarker(data)
	fmt.Printf("Part 1 First Marker: %v\n", firstMarker)
}

func readFile() string {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var dataStream string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		dataStream = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return dataStream
}

func findFirstMarker(data string) int {
	for i := 4; i < len(data); i++ {
		set := append([]string{}, strings.Split(data[i-4:i], "")...)
		sort.Strings(set)
		for j, character := range set {
			if j == 0 {
				continue
			}
			if set[j-1] == character {
				break
			}
			if j == 3 {
				return i
			}
		}
	}

	return 0
}
