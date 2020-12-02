package main

import (
	"fmt"
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
	fmt.Printf("Test case 3 is true, %v", isValidPassword(data{
		minimum:  2,
		maximum:  9,
		letter:   "c",
		password: "ccccccccc",
	}))
}

func isValidPassword(data data) bool {
	count := strings.Count(data.password, data.letter)
	if count >= data.minimum && count <= data.maximum {
		return true
	}

	return false
}
