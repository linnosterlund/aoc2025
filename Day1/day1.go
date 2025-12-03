package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	startPosition := 50
	lines := readFile()

	p1 := part1(startPosition, lines)
	p2 := part2(startPosition, lines)

	fmt.Printf("The password in the first part is: %d\n", p1)
	fmt.Printf("The password in the second part is: %d\n", p2)
}

func readFile() []string {
	var lines []string

	file, err := os.Open("puzzleInput.txt")
	if err != nil {
		log.Fatal(err)
		return lines
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
		return lines
	}

	return lines
}

func part2(currentPosition int, lines []string) int {
	password := 0

	for _, line := range lines {
		var direction rune
		var distance int
		fmt.Sscanf(line, "%c%d", &direction, &distance)

		oldPosition := currentPosition

		password += distance / 100
		distance = distance - (distance / 100 * 100)

		switch direction {
		case 'L':
			currentPosition -= distance
		case 'R':
			currentPosition += distance
		}

		switch {
		case currentPosition == 0 && distance != 0:
			password++
		case currentPosition < 0 && oldPosition > 0:
			password++
		case currentPosition > 99:
			password++
		}

		currentPosition = (currentPosition + 100) % 100
	}

	return password
}

func part1(currentPosition int, lines []string) int {
	password := 0

	for _, line := range lines {
		var direction rune
		var distance int
		fmt.Sscanf(line, "%c%d", &direction, &distance)

		switch direction {
		case 'L':
			currentPosition = (currentPosition - distance + 100) % 100
		case 'R':
			currentPosition = (currentPosition + distance) % 100
		}

		if currentPosition == 0 {
			password++
		}
	}

	return password
}
