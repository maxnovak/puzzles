package main

import (
	"fmt"
	"strconv"
)

func main() {
	//Test Cases
	fmt.Println("111111 ", checkIfValid("111111"))
	fmt.Println("223450 ", checkIfValid("223450"))
	fmt.Println("123789 ", checkIfValid("123789"))
	fmt.Println("112233 ", checkIfValid("112233"))
	fmt.Println("123444 ", checkIfValid("123444"))
	fmt.Println("111122 ", checkIfValid("111122"))

	validPasswords := 0
	validPart2Passwords := 0

	for password := 138307; password < 654504; password++ {
		passwordString := strconv.Itoa(password)
		if (checkIfValid(passwordString)) {
			validPasswords++
		}
		if (checkIfValid2(passwordString)) {
			validPart2Passwords++
		}
	}

	fmt.Println(validPasswords)
	fmt.Println(validPart2Passwords)
}

func checkIfValid(password string) bool{
	mightBeValid := false
	for i :=0; i < len(password) -1; i++ {
		firstChar := string(password[i])
		secondChar := string(password[i+1])

		if (firstChar > secondChar) {
			return false
		}
		if (firstChar == secondChar) {
			mightBeValid = true
		}
	}
	return mightBeValid
}

func checkIfValid2(password string) bool{
	mightBeValid := false
	currentRun := 1
	for i :=0; i < len(password) -1; i++ {
		firstChar := string(password[i])
		secondChar := string(password[i+1])

		if (firstChar > secondChar) {
			return false
		}
		if (firstChar == secondChar) {
			currentRun++
		} else {
			if currentRun == 2 {
				mightBeValid = true
			}
			currentRun = 1
		}
	}

	if (currentRun == 2) {
		mightBeValid = true
	}

	return mightBeValid
}