package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("this is day 6")
	forms := readFile()
	sum := 0
	for _, item := range forms {
		sum += len(item)
	}
	fmt.Println(sum)
}

func readFile() []map[string]bool {

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var customsForms []map[string]bool
	groupsForm := map[string]bool{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			customsForms = append(customsForms, groupsForm)
			groupsForm = map[string]bool{}
			continue
		}
		for _, item := range text {
			groupsForm[string(item)] = true
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return customsForms
}
