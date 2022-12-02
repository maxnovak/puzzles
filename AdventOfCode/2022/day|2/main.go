package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Data struct {
	YourMove  int
	TheirMove int
	Score     int
}

func main() {
	fmt.Println("This is day 2")
	tournament := readFile()
	score := resultWithMatches(tournament)

	fmt.Printf("Total Score Part 1: %v \n", score)
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
			YourMove:  convertMoveToInt(output[1]),
			TheirMove: convertMoveToInt(output[0]),
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
