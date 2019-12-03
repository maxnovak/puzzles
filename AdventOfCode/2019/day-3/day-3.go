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
	coord1, coord2 := readFile()
	fmt.Println(coord1)
	fmt.Println(coord2)
	graph1 := graph(coord1)
	graph2 := graph(coord2)
	fmt.Println(graph1)
	fmt.Println(graph2)
	findMatches(graph1, graph2)
}


func findMatches(graph1 [][]int, graph2 [][]int) {
	for _, e1 := range graph1 {
		for _, e2 := range graph2 {
			if e1[0] == e2[0] && e1[1] == e2[1] {
				fmt.Println(e1)
			}
		}
	}
}


func graph(coord1 []string) [][]int{
	var traversal [][]int
	currentx := 1
	currenty := 1
	traversal = append(traversal, []int{currentx, currenty})
	for _, location := range coord1 {
		distance, _ := strconv.Atoi(location[1:len(location)])
		var newx int
		var newy int
		if location[0] == 'R' {
			distance = distance  + currentx + 1
			for x := currentx+1; x < distance; x++ {
				traversal = append(traversal, []int{x, currenty})
				newx = x
			}
			currentx = newx
			fmt.Println("Right", currentx)
		} else if location[0] == 'L' {
			distance = currentx - distance -1
			for x := currentx-1; x > distance; x-- {
				traversal = append(traversal, []int{x, currenty})
				newx = x
			}
			currentx = newx
			fmt.Println("Left", currentx, distance)
		} else if location[0] == 'U' {
			distance = distance  + currenty + 1
			for y := currenty+1; y < distance; y++ {
				traversal = append(traversal, []int{currentx, y})
				newy = y
			}
			currenty = newy
			fmt.Println("Up", location[1:len(location)])
		} else if location[0] == 'D' {
			distance = currenty - distance -1
			for y := currenty-1; y > distance; y-- {
				traversal = append(traversal, []int{currentx, y})
				newy = y
			}
			currenty = newy
			fmt.Println("Down", location[1:len(location)])
		}
	}
	return traversal
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