package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type data struct {
	minimum  int
	maximum  int
	letter   string
	password string
}

func main() {
	fmt.Println("This is day 2")
	fmt.Printf("Test case 1 is true, %v\n", isValidPassword(data{
		minimum:  1,
		maximum:  3,
		letter:   "a",
		password: "abcde",
	}))
	fmt.Printf("Test case 2 is false, %v\n", isValidPassword(data{
		minimum:  1,
		maximum:  3,
		letter:   "b",
		password: "cdefg",
	}))
	fmt.Printf("Test case 3 is true, %v\n", isValidPassword(data{
		minimum:  2,
		maximum:  9,
		letter:   "c",
		password: "ccccccccc",
	}))

	passwords := readFile()
	count := 0
	for _, pass := range passwords {
		if isValidPassword(pass) {
			count++
		}
	}
	fmt.Printf("Valid number of passwords: %v\n", count)

	fmt.Printf("Test case 1 is true, %v\n", isValidPasswordV2(data{
		minimum:  1,
		maximum:  3,
		letter:   "a",
		password: "abcde",
	}))
	fmt.Printf("Test case 2 is false, %v\n", isValidPasswordV2(data{
		minimum:  1,
		maximum:  3,
		letter:   "b",
		password: "cdefg",
	}))
	fmt.Printf("Test case 3 is false, %v\n", isValidPasswordV2(data{
		minimum:  2,
		maximum:  9,
		letter:   "c",
		password: "ccccccccc",
	}))

	count2 := 0
	for _, pass := range passwords {
		if isValidPasswordV2(pass) {
			count2++
		}
	}
	fmt.Printf("Valid number of passwords: %v\n", count2)

}

func isValidPassword(data data) bool {
	count := strings.Count(data.password, data.letter)
	if count >= data.minimum && count <= data.maximum {
		return true
	}

	return false
}

func isValidPasswordV2(data data) bool {
	if (string(data.password[data.minimum-1]) == data.letter && string(data.password[data.maximum-1]) != data.letter) || (string(data.password[data.minimum-1]) != data.letter && string(data.password[data.maximum-1]) == data.letter) {
		return true
	}
	return false
}

func readFile() []data {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passwords []data
	for scanner.Scan() {
		output := strings.Split(scanner.Text(), " ")
		ranges := strings.Split(output[0], "-")
		minimum, _ := strconv.Atoi(ranges[0])
		maximum, _ := strconv.Atoi(ranges[1])

		passwords = append(passwords, data{
			minimum:  minimum,
			maximum:  maximum,
			letter:   string(output[1][0]),
			password: output[2],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passwords
}
