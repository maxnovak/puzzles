package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Test Cases
	fmt.Println("111111 ", checkIfValid("111111"))
	fmt.Println("223450, ", checkIfValid("223450"))
	fmt.Println("123789 ", checkIfValid("123789"))

}

func checkIfValid(password string) bool{
	mightBeValid := false
	for i :=0; i < len(password) -1; i++ {
		firstChar, _ := strconv.Atoi(string(password[i]))
		secondChar, _ := strconv.Atoi(string(password[i+1]))

		if (firstChar > secondChar) {
			return false
		}
		if (firstChar == secondChar) {
			mightBeValid = true
		}
	}
	return mightBeValid
}