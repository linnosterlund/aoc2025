package main

/*
- every battery has a joltage rating 1-9
- one bank: a line of digits in the  input
- for each bank turn on two batteries
	a. joltage is equal to the number formed by the digits turned on
	b. you cannot rearrange batteries to form another digit order
	c. find the largest joltage for each bank
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	totalJoltage := 0
	banks := getBanks()
	for _, bank := range banks {
		number := checkHighestPair(bank)
		totalJoltage += number
	}

	fmt.Printf("Part 1: total joltage is: %d\n", totalJoltage)

	totalJoltage = 0

	for _, bank := range banks {
		number := checkHighestTwelve(bank)
		totalJoltage += number
	}

	fmt.Printf("Part 2: total joltage is: %d\n", totalJoltage)
}

func checkHighestTwelve(bank string) int {
	highestDigit := getInt(bank[0])
	startIndex := 1
	highestIndex := 0
	lastIndex := len(bank) - 11
	var digits []int

	for range 12 {
		for i := startIndex; i < lastIndex; i++ {
			digit := getInt(bank[i])

			if digit > highestDigit {
				highestDigit = digit
				highestIndex = i
			}
		}

		digits = append(digits, highestDigit)
		startIndex = highestIndex + 2
		highestIndex++
		lastIndex++
		if highestIndex < len(bank) {
			highestDigit = getInt(bank[highestIndex])
		}
	}

	number := sliceToInt(digits)
	return number
}

func checkHighestPair(bank string) int {
	length := len(bank)

	firstDigit := getInt(bank[0])
	firstDigitIndex := 0

	for i := 1; i < length-1; i++ {
		digit := getInt(bank[i])

		if digit > firstDigit {
			firstDigit = digit
			firstDigitIndex = i
		}
	}

	secondDigit := getInt(bank[firstDigitIndex+1])
	for i := firstDigitIndex + 2; i < length; i++ {
		digit := getInt(bank[i])

		if digit > secondDigit {
			secondDigit = digit
		}
	}

	number := (firstDigit * 10) + secondDigit
	return number
}

func getInt(b byte) int {
	n, err := strconv.Atoi(string(b))
	if err != nil {
		fmt.Printf("error parsing byte %b", b)
	}
	return n
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}

func getBanks() []string {
	var banks []string

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error reading input", err)
		return banks
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		bank := scanner.Text()
		banks = append(banks, bank)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return banks
	}

	return banks
}
