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
	graph(coord1)
	graph(coord2)
}

func graph(coord1 []string) {
	for _, location := range coord1 {
		if location[0] == 'R' {
			fmt.Println("Right", location[1:len(location)])
		} else if location[0] == 'L' {
			fmt.Println("Left", location[1:len(location)])
		} else if location[0] == 'U' {
			fmt.Println("Up", location[1:len(location)])
		} else if location[0] == 'D' {
			fmt.Println("Down", location[1:len(location)])
		}
	}
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