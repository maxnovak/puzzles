package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Data struct {
	RangeOneStart int
	RangeOneEnd   int
	RangeTwoStart int
	RangeTwoEnd   int
}

func main() {
	fmt.Println("This is day 4")
	pairsRanges := readFile()
	count := 0
	count2 := 0
	for _, pair := range pairsRanges {
		if pair.RangeOneStart <= pair.RangeTwoStart &&
			pair.RangeTwoEnd <= pair.RangeOneEnd {
			count++
		} else if pair.RangeTwoStart <= pair.RangeOneStart &&
			pair.RangeOneEnd <= pair.RangeTwoEnd {
			count++
		}

		if pair.RangeOneStart <= pair.RangeTwoStart &&
			pair.RangeTwoStart <= pair.RangeOneEnd {
			count2++
		} else if pair.RangeTwoStart <= pair.RangeOneStart &&
			pair.RangeOneStart <= pair.RangeTwoEnd {
			count2++
		}
	}
	fmt.Printf("Count for fully contains: %v \n", count)
	fmt.Printf("Count for overlap: %v \n", count2)
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
		pairs := strings.Split(scanner.Text(), ",")
		first := strings.Split(pairs[0], "-")
		second := strings.Split(pairs[1], "-")
		rangeOneStart, _ := strconv.Atoi(first[0])
		rangeOneEnd, _ := strconv.Atoi(first[1])
		rangeTwoStart, _ := strconv.Atoi(second[0])
		rangeTwoEnd, _ := strconv.Atoi(second[1])
		data = append(data, Data{
			RangeOneStart: rangeOneStart,
			RangeOneEnd:   rangeOneEnd,
			RangeTwoStart: rangeTwoStart,
			RangeTwoEnd:   rangeTwoEnd,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return data
}
