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

	fmt.Printf("%v", intcodes)

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