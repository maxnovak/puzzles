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
	fmt.Println("Day 8")
	forest := readFile()
	total, highestScenicScore := findVisible(forest)
	fmt.Printf("Part 1 Total: %v\n", total)
	fmt.Printf("Part 2: Highest Score: %v\n", highestScenicScore)
}

func readFile() [][]int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	data := [][]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		rowData := strings.Split(scanner.Text(), "")
		intData := []int{}
		for _, value := range rowData {
			parsed, _ := strconv.Atoi(value)
			intData = append(intData, parsed)
		}
		data = append(data, intData)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func findVisible(forest [][]int) (int, int) {
	total := 0
	highestScenicScore := 0
	for i, row := range forest {
		for j, tree := range row {
			if j == 0 || i == 0 || j == len(row)-1 || i == len(forest)-1 {
				total++
				continue
			}
			height := tree
			visibleLeft := true
			leftScore := 0
			visibleRight := true
			rightScore := 0
			visibleTop := true
			topScore := 0
			visibleBottom := true
			bottomScore := 0
			for k := j - 1; k >= 0; k-- { //check left
				leftScore++
				if height <= row[k] {
					visibleLeft = false
					break
				}
			}
			for k := j + 1; k < len(row); k++ { //check right
				rightScore++
				if height <= row[k] {
					visibleRight = false
					break
				}
			}
			for k := i - 1; k >= 0; k-- { //check top
				topScore++
				if height <= forest[k][j] {
					visibleTop = false
					break
				}
			}
			for k := i + 1; k < len(forest); k++ { //check right
				bottomScore++
				if height <= forest[k][j] {
					visibleBottom = false
					break
				}
			}
			score := leftScore * rightScore * topScore * bottomScore
			if score > highestScenicScore {
				highestScenicScore = score
			}
			if visibleBottom || visibleLeft || visibleRight || visibleTop {
				total++
			}
		}
	}

	return total, highestScenicScore
}
