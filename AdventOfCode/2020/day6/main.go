package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type groupsForms struct {
	total           int
	questionAnswers map[string]int
}

func main() {
	fmt.Println("this is day 6")
	forms := readFile()
	part1Sum := 0
	part2Sum := 0
	for _, item := range forms {
		for _, answers := range item.questionAnswers {
			if item.total == answers {
				part2Sum++
			}
			part1Sum++
		}
	}
	fmt.Println("Part 1 solution:", part1Sum, "\nPart 2 solution:", part2Sum)
}

func readFile() []groupsForms {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var customsForms []groupsForms
	groupsForm := groupsForms{
		total:           0,
		questionAnswers: map[string]int{},
	}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			customsForms = append(customsForms, groupsForm)
			groupsForm = groupsForms{
				total:           0,
				questionAnswers: map[string]int{},
			}
			continue
		}

		for _, item := range text {
			groupsForm.questionAnswers[string(item)]++
		}
		groupsForm.total++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return customsForms
}
