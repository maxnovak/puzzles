package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type data struct {
	elfNumber     int
	totalCalories int
}

func main() {
	fmt.Println("This is day 1")
	data := readFile()

	maxCalories := 0
	for index, item := range data {
		if index == 0 || item.totalCalories > maxCalories {
			maxCalories = item.totalCalories
		}
	}

	fmt.Printf("Most calories: %v", maxCalories)
}

func readFile() []data {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var elfList []data
	totalCal := 0
	elfNumber := 1
	for scanner.Scan() {
		output := scanner.Text()
		if output == "" {
			elfList = append(elfList, data{
				elfNumber:     elfNumber,
				totalCalories: totalCal,
			})
			elfNumber++
			totalCal = 0
		} else {
			parsedValue, _ := strconv.Atoi(output)
			totalCal += parsedValue
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return elfList
}
