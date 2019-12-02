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

	for i := 0; i < len(intcodes); i++ {
		intcode := intcodes[i]
		if (intcode == 99) {
			fmt.Println("done")
			break
		}
		if (intcode == 1) {
		fmt.Println(i)
			valueLocation1 := intcodes[i+1]
			valueLocation2 := intcodes[i+2]
			storeLocation := intcodes[i+3]
			intcodes[storeLocation] = intcodes[valueLocation1] + intcodes[valueLocation2]
			i = i+3
		}
		if (intcode == 2) {
			fmt.Println(i)
			valueLocation1 := intcodes[i+1]
			valueLocation2 := intcodes[i+2]
			storeLocation := intcodes[i+3]
			intcodes[storeLocation] = intcodes[valueLocation1] * intcodes[valueLocation2]
			i = i+3
		}
	}
	fmt.Println(intcodes)
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