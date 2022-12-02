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
	score := result(tournament)

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
			YourMove:  convertStringToInt(output[1]),
			TheirMove: convertStringToInt(output[0]),
		})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func result(listOfMatches []Data) int {
	totalScore := 0
	for i, match := range listOfMatches {
		if match.YourMove == match.TheirMove {
			listOfMatches[i].Score = match.YourMove + 3
			totalScore += match.YourMove + 3
			continue
		}
		if match.TheirMove == 1 && match.YourMove == 3 {
			listOfMatches[i].Score = match.YourMove
			totalScore += match.YourMove
			continue
		}
		if match.YourMove < match.TheirMove && !(match.YourMove == 1 && match.TheirMove == 3) {
			listOfMatches[i].Score = match.YourMove
			totalScore += match.YourMove
			continue
		}
		listOfMatches[i].Score = match.YourMove + 6
		totalScore += match.YourMove + 6
	}
	return totalScore
}

func convertStringToInt(input string) int {
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
