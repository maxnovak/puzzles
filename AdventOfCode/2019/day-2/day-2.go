package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	intcodes := readFile()

	intcodes[1] = 12
	intcodes[2] = 2

	fmt.Println(doMath(intcodes)) // Solution for Part 1 5305097

	for value1 := 0; value1 < 100; value1++ {
		for value2 := 0; value2 < 100; value2++ {
			intcodes = readFile()
			intcodes[1] = value1
			intcodes[2] = value2
			result := doMath(intcodes)
			if (result[0] == 19690720) {
				fmt.Println(value1 * 100 + value2)
				break
			}
		}
	}
}

func doMath(intcodes []int) []int{
	i := 0
	for true {
		optcode := intcodes[i]
		if (optcode == 99) {
			break
		}
		valueLocation1 := intcodes[i+1]
		valueLocation2 := intcodes[i+2]
		storeLocation := intcodes[i+3]
		if (optcode == 1) {
			intcodes[storeLocation] = intcodes[valueLocation1] + intcodes[valueLocation2]
		}
		if (optcode == 2) {
			intcodes[storeLocation] = intcodes[valueLocation1] * intcodes[valueLocation2]
		}
		i += 4
	}
	return intcodes
}

func readFile() []int{

	file, err := os.Open("day-2-input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var intcodes []int
    for scanner.Scan() {
    	input := strings.Split(scanner.Text(), ",")
    	for _, i := range input {
			number, _ := strconv.Atoi(i)
        	intcodes = append(intcodes, number)
    	}
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return intcodes
}