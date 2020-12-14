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
	location := 0
	visitedLocations := map[int]bool{}
	accumulatedValue := 0
	for {
		if visitedLocations[location] {
			break
		}
		visitedLocations[location] = true
		op := ops[location]
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
	fmt.Println("Part 1 solution", accumulatedValue)
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
