package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fmt.Println("This is day 3")
	passports := readFile()

	count := 0
	for _, passport := range passports {
		valid := checkValidPassport(passport)
		if valid {
			count++
		}
	}
	fmt.Printf("Number of valid passport like things: %v", count)
}

func checkValidPassport(passport map[string]string) bool {
	if passport["byr"] != "" && passport["iyr"] != "" && passport["eyr"] != "" && passport["hgt"] != "" && passport["hcl"] != "" && passport["ecl"] != "" && passport["pid"] != "" {
		return true
	}
	return false
}

func readFile() []map[string]string {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var passports []map[string]string
	passport := map[string]string{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			passports = append(passports, passport)
			passport = map[string]string{}
			continue
		}
		line := strings.Split(text, " ")
		fmt.Printf("%v, %v, %v,", len(line), line, text)
		for _, item := range line {
			keyValue := strings.Split(item, ":")
			passport[keyValue[0]] = keyValue[1]
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return passports
}
