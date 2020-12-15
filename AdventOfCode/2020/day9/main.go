package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("this is day 9")
	numbers := readFile()
	locationOfFailure := 0
	for i := 25; i < len(numbers); i++ {
		if !findIfSum(numbers[i-25:i], numbers[i]) {
			fmt.Println("Part 1 solution", numbers[i])
			locationOfFailure = i
			break
		}
	}
	for i := 0; i < locationOfFailure; i++ {
		for j := i + 1; j < locationOfFailure; j++ {
			sum, max, min := sumSlice(numbers[i:j])
			if sum == numbers[locationOfFailure] {
				fmt.Println("Part 2 solution", max+min)
			}
		}
	}
}

func sumSlice(numbers []int) (int, int, int) {
	result := 0
	max := 0
	min := numbers[0]
	for _, number := range numbers {
		if number > max {
			max = number
		}
		if number < min {
			min = number
		}
		result += number
	}
	return result, max, min
}

func findIfSum(numbers []int, sum int) bool {
	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i]+numbers[j] == sum {
				return true
			}
		}
	}
	return false
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
