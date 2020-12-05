package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type boardingPass struct {
	row    string
	column string
}

func main() {
	fmt.Println("this is day 5")
	input := readFile()
	seatID := 0
	for _, boardingPass := range input {
		row := findRow(boardingPass.row)
		column := findColumn(boardingPass.column)
		testSeatID := (row * 8) + column
		if testSeatID > seatID {
			seatID = testSeatID
		}
	}
	fmt.Printf("highest Seat ID: %v", seatID)
}

func findRow(row string) int {
	rows := []int{}
	for i := 0; i < 128; i++ {
		rows = append(rows, i)
	}
	for _, item := range row {
		if item == 'B' {
			rows = rows[len(rows)/2:]
		}
		if item == 'F' {
			rows = rows[0 : len(rows)/2]
		}
	}
	return rows[0]
}

func findColumn(column string) int {
	columns := []int{0, 1, 2, 3, 4, 5, 6, 7}
	for _, item := range column {
		if item == 'R' {
			columns = columns[len(columns)/2:]
		}
		if item == 'L' {
			columns = columns[0 : len(columns)/2]
		}
	}
	return columns[0]
}

func readFile() []boardingPass {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var boardingPasses []boardingPass
	for scanner.Scan() {
		pass := scanner.Text()
		boardingPasses = append(boardingPasses, boardingPass{
			row:    pass[0:7],
			column: pass[7:10],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return boardingPasses
}
