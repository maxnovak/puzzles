package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("This is day 1")
	data := readFile()
	sort.Ints(data)

	fmt.Printf("Most calories: %v", data[len(data)-1])
	fmt.Printf("Second most calories: %v", data[len(data)-2])
	fmt.Printf("Third most calories: %v", data[len(data)-3])
	fmt.Printf("Sum of calories: %v", data[len(data)-1]+data[len(data)-2]+data[len(data)-3])

}

func readFile() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var elfList []int
	totalCal := 0
	for scanner.Scan() {
		output := scanner.Text()
		if output == "" {
			elfList = append(elfList, totalCal)
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
