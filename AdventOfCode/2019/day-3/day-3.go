package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	coord1, coord2 := readFile()
	fmt.Println(coord1)
	fmt.Println(coord2)
}

func readFile() ([]string, []string) {

	file, err := os.Open("day-3-input.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var coordinates1 []string
    var coordinates2 []string
    ind := 0
    for scanner.Scan() {
    	if (ind == 0) {
	    	directions := strings.Split(scanner.Text(), ",")
	    	for _, i := range directions {
	        	coordinates1 = append(coordinates1, i)
	    	}
	    }
	    if (ind == 1) {
	    	directions := strings.Split(scanner.Text(), ",")
	    	for _, i := range directions {
	        	coordinates2 = append(coordinates2, i)
	    	}	
	    }
	    ind++
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    return coordinates1, coordinates2
}