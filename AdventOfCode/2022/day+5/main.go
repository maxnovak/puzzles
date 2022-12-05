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
	fmt.Println("This is day 5")
	crates, movement := readFile()
	crateState := executeMovement(crates, movement)
	for _, column := range crateState {
		fmt.Printf(column[0])
	}
}

type Data struct {
	numberToMove int
	startColumn  int
	endColumn    int
}

func readFile() ([][]string, []Data) {
	file, err := os.Open("input_crates.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	data := make([][]string, 9)
	for scanner.Scan() {
		row := scanner.Text()
		if row[1] != ' ' {
			newValue := []string{string(row[1])}
			data[0] = append(data[0], newValue...)
		}
		if row[5] != ' ' {
			newValue := []string{string(row[5])}
			data[1] = append(data[1], newValue...)
		}
		if row[9] != ' ' {
			newValue := []string{string(row[9])}
			data[2] = append(data[2], newValue...)
		}
		if row[13] != ' ' {
			newValue := []string{string(row[13])}
			data[3] = append(data[3], newValue...)
		}
		if row[17] != ' ' {
			newValue := []string{string(row[17])}
			data[4] = append(data[4], newValue...)
		}
		if row[21] != ' ' {
			newValue := []string{string(row[21])}
			data[5] = append(data[5], newValue...)
		}
		if row[25] != ' ' {
			newValue := []string{string(row[25])}
			data[6] = append(data[6], newValue...)
		}
		if row[29] != ' ' {
			newValue := []string{string(row[29])}
			data[7] = append(data[7], newValue...)
		}
		if row[33] != ' ' {
			newValue := []string{string(row[33])}
			data[8] = append(data[8], newValue...)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	file, err = os.Open("input_movement.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner = bufio.NewScanner(file)
	var dataMovement []Data
	for scanner.Scan() {
		commands := strings.Split(scanner.Text(), " ")
		numberToMove, _ := strconv.Atoi(commands[1])
		start, _ := strconv.Atoi(commands[3])
		end, _ := strconv.Atoi(commands[5])
		dataMovement = append(dataMovement, Data{
			numberToMove: numberToMove,
			startColumn:  start,
			endColumn:    end,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data, dataMovement
}

func executeMovement(crates [][]string, commands []Data) [][]string {
	fmt.Println(crates)
	for _, command := range commands {
		for i := 0; i < command.numberToMove; i++ {
			moving := crates[command.startColumn-1][:1]
			newStartCol := append([]string{}, crates[command.startColumn-1][1:]...)
			newEndCol := append(moving, crates[command.endColumn-1]...)
			crates[command.endColumn-1] = newEndCol
			crates[command.startColumn-1] = newStartCol
		}
	}

	return crates
}
