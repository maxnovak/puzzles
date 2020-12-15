package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Println("this is day 10")

	adapters := append(readFile(), 0)
	sort.Ints(adapters)
	threes := 1
	ones := 0
	for i := 0; i < len(adapters)-1; i++ {
		diff := adapters[i+1] - adapters[i]
		fmt.Println(diff, adapters[i], adapters[i+1])
		if diff == 1 {
			ones++
		}
		if diff == 3 {
			threes++
		}
	}
	fmt.Println("threes:", threes, "ones:", ones)
	fmt.Println("Part 1 solution:", threes*ones)
}

func readFile() []int {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		text := scanner.Text()
		number, _ := strconv.Atoi(text)
		numbers = append(numbers, number)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return numbers
}
