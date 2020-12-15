package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type operation struct {
	name   string
	amount int
}

func main() {
	fmt.Println("this is day 8")
	ops := readFile()

	accumulatedValue, _ := findLoopPoint(ops)

	fmt.Println("Part 1 solution", accumulatedValue)

	for i := 0; i < len(ops); i++ {
		if ops[i].name == "nop" {
			ops[i].name = "jmp"
			accumulatedValue, foundIt := findLoopPoint(ops)
			ops[i].name = "nop"
			if foundIt {
				fmt.Println("Part 2 solution:", accumulatedValue)
				break
			}
		} else if ops[i].name == "jmp" {
			ops[i].name = "nop"
			accumulatedValue, foundIt := findLoopPoint(ops)
			ops[i].name = "jmp"
			if foundIt {
				fmt.Println("Part 2 solution:", accumulatedValue)
				break
			}
		}
	}
}

func findLoopPoint(operations []operation) (int, bool) {
	location := 0
	visitedLocations := map[int]bool{}
	accumulatedValue := 0
	foundIt := false
	for {
		if visitedLocations[location] {
			break
		}
		if location == len(operations) {
			foundIt = true
			break
		}
		visitedLocations[location] = true
		op := operations[location]
		if op.name == "nop" {
			location++
			continue
		}
		if op.name == "acc" {
			accumulatedValue += op.amount
			location++
			continue
		}
		if op.name == "jmp" {
			location += op.amount
			continue
		}
	}
	return accumulatedValue, foundIt
}

func readFile() []operation {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var operations []operation
	for scanner.Scan() {
		text := scanner.Text()
		text = strings.Replace(text, "+", "", -1)
		line := strings.Split(text, " ")
		amount, _ := strconv.Atoi(line[1])
		operations = append(operations, operation{
			name:   line[0],
			amount: amount,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return operations
}
