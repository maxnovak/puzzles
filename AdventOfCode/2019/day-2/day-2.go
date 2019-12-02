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

	i := 0
	intcodes[1] = 12
	intcodes[2] = 2
	for true {
		optcode := intcodes[i]
		if (optcode == 99) {
			fmt.Println("done")
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