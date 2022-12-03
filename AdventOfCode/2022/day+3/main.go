package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Data struct {
	fullBag        string
	compartmentOne string
	compartmentTwo string
	sharedItem     rune
	priority       int
}

var characterValues = map[rune]int{
	'a': 1,
	'b': 2,
	'c': 3,
	'd': 4,
	'e': 5,
	'f': 6,
	'g': 7,
	'h': 8,
	'i': 9,
	'j': 10,
	'k': 11,
	'l': 12,
	'm': 13,
	'n': 14,
	'o': 15,
	'p': 16,
	'q': 17,
	'r': 18,
	's': 19,
	't': 20,
	'u': 21,
	'v': 22,
	'w': 23,
	'x': 24,
	'y': 25,
	'z': 26,
	'A': 27,
	'B': 28,
	'C': 29,
	'D': 30,
	'E': 31,
	'F': 32,
	'G': 33,
	'H': 34,
	'I': 35,
	'J': 36,
	'K': 37,
	'L': 38,
	'M': 39,
	'N': 40,
	'O': 41,
	'P': 42,
	'Q': 43,
	'R': 44,
	'S': 45,
	'T': 46,
	'U': 47,
	'V': 48,
	'W': 49,
	'X': 50,
	'Y': 51,
	'Z': 52,
}

func main() {
	fmt.Println("This is day 3")
	bagContents := readFile()
	totalPriority := 0
	totalPriority2 := 0
	for i, item := range bagContents {
		bagContents[i].sharedItem = findSharedItem(item.compartmentOne, item.compartmentTwo)
		bagContents[i].priority = characterValues[bagContents[i].sharedItem]
		totalPriority += bagContents[i].priority
	}

	for i := 0; i < len(bagContents); i += 3 {
		character := findSharedItems(bagContents[i].fullBag, bagContents[i+1].fullBag, bagContents[i+2].fullBag)
		totalPriority2 += characterValues[character]
	}

	fmt.Printf("Total Priority Part 1: %v \n", totalPriority)
	fmt.Printf("Total Priority Part 2: %v \n", totalPriority2)

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
		output := scanner.Text()
		length := len(output)
		data = append(data, Data{
			fullBag:        output,
			compartmentOne: output[:length/2],
			compartmentTwo: output[length/2:],
		})

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}

func findSharedItem(listOne string, listTwo string) rune {
	var sharedItem rune
	for _, item := range listOne {
		for _, compareItem := range listTwo {
			if item == compareItem {
				return item
			}
		}
	}
	return sharedItem
}

func findSharedItems(listOne string, listTwo string, listThree string) rune {
	var sharedItem rune
	for _, item := range listOne {
		for _, compareItem := range listTwo {
			for _, compareItem2 := range listThree {
				if item == compareItem && item == compareItem2 {
					return item
				}
			}
		}
	}
	return sharedItem
}
