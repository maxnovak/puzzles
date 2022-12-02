package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Data struct {
	YourMove     int
	TheirMove    int
	NeededResult int
	Score        int
}

func main() {
	fmt.Println("This is day 2")
	tournament := readFile()
	score := resultWithMatches(tournament)
	score2 := resultWithResult(tournament)

	fmt.Printf("Total Score Part 1: %v \n", score)
	fmt.Printf("Total Score Part 2: %v \n", score2)
}

func readFile() []Data {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var data []Data
	for scanner.Scan() {
		output := strings.Split(scanner.Text(), " ")
		data = append(data, Data{
			YourMove:     convertMoveToInt(output[1]),
			TheirMove:    convertMoveToInt(output[0]),
			NeededResult: convertResultToInt(output[1]),
		})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func resultWithMatches(listOfMatches []Data) int {
	totalScore := 0
	for i, match := range listOfMatches {
		result := determineWinOrLossOrDraw(match.TheirMove, match.YourMove)
		listOfMatches[i].Score = match.YourMove + result
		totalScore += match.YourMove + result
	}
	return totalScore
}

func resultWithResult(listOfMatches []Data) int {
	totalScore := 0
	for _, match := range listOfMatches {
		if match.NeededResult == 3 {
			totalScore += match.TheirMove + 3
			continue
		}
		for i := 1; i < 4; i++ {
			if determineWinOrLossOrDraw(match.TheirMove, i) == match.NeededResult {
				totalScore += match.NeededResult + i
				fmt.Printf("Result: %v, Move: %v Their Move: %v\n ", match.NeededResult, i, match.TheirMove)
			}
		}
	}
	return totalScore
}

func determineWinOrLossOrDraw(them int, you int) int {
	if you == them {
		return 3
	}
	if them == 1 && you == 3 {
		return 0
	}
	if you < them && !(you == 1 && them == 3) {
		return 0
	}
	return 6
}

func convertResultToInt(input string) int {
	if input == "X" {
		return 0
	}
	if input == "Y" {
		return 3
	}
	if input == "Z" {
		return 6
	}
	return 0
}

func convertMoveToInt(input string) int {
	if input == "A" || input == "X" {
		return 1 //Rock
	}
	if input == "B" || input == "Y" {
		return 2 //Paper
	}
	if input == "C" || input == "Z" {
		return 3 //Scissors
	}
	return 0 //Error
}
