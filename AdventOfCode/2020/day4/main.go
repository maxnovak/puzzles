package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("This is day 3")
	passports := readFile()

	countSimple := 0
	countComplex := 0
	for _, passport := range passports {
		validSimple := checkValidPassport(passport)
		if validSimple {
			countSimple++
		}
		validComplex := checkValidPassportDetailed(passport)
		if validComplex {
			countComplex++
		}
	}
	fmt.Printf("Number of valid passport like things simple: %v complex: %v", countSimple, countComplex)
}

func checkValidPassport(passport map[string]string) bool {
	if passport["byr"] != "" && passport["iyr"] != "" && passport["eyr"] != "" && passport["hgt"] != "" && passport["hcl"] != "" && passport["ecl"] != "" && passport["pid"] != "" {
		return true
	}
	return false
}

func checkValidPassportDetailed(passport map[string]string) bool {
	valid := false

	if passport["byr"] != "" && passport["iyr"] != "" && passport["eyr"] != "" && passport["hgt"] != "" && passport["hcl"] != "" && passport["ecl"] != "" && passport["pid"] != "" {
		birthYear, _ := strconv.Atoi(passport["byr"])
		issueYear, _ := strconv.Atoi(passport["iyr"])
		expirationYear, _ := strconv.Atoi(passport["eyr"])
		heightAmount, _ := strconv.Atoi(passport["hgt"][0 : len(passport["hgt"])-2])
		heightUnit := passport["hgt"][len(passport["hgt"])-2:]

		if birthYear <= 2002 && birthYear >= 1920 {
			valid = true
		} else {
			return false
		}
		if issueYear <= 2020 && issueYear >= 2010 {
			valid = true
		} else {
			return false
		}
		if expirationYear <= 2030 && expirationYear >= 2020 {
			valid = true
		} else {
			return false
		}
		if heightUnit == "cm" && heightAmount <= 193 && heightAmount >= 150 {
			valid = true
		} else if heightUnit == "in" && heightAmount <= 76 && heightAmount >= 59 {
			valid = true
		} else {
			return false
		}
		hclresult, _ := regexp.MatchString(`#[0-9a-z]{6}`, passport["hcl"])
		if hclresult && len(passport["hcl"]) == 7 {
			valid = true
		} else {
			return false
		}
		if passport["ecl"] == "amb" || passport["ecl"] == "blu" || passport["ecl"] == "brn" || passport["ecl"] == "gry" || passport["ecl"] == "grn" || passport["ecl"] == "hzl" || passport["ecl"] == "oth" {
			valid = true
		} else {
			return false
		}
		pidResult, _ := regexp.MatchString(`[0-9]{9}`, passport["pid"])
		if len(passport["pid"]) == 9 && pidResult {
			valid = true
		} else {
			return false
		}
	}
	return valid
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
