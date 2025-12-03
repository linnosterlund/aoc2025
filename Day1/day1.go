package main

/***
Day 1: Find the password

"Due to new security protocols, the password is locked in the safe below. Please see the attached document for the new combination."

The safe has a dial with only an arrow on it; around the dial are the numbers 0 through 99 in order. As you turn the dial, it makes a small click noise as it reaches each number.

The attached document (your puzzle input) contains a sequence of rotations, one per line, which tell you how to open the safe. A rotation starts with an L or R which indicates whether the rotation should be to the left (toward lower numbers) or to the right (toward higher numbers). Then, the rotation has a distance value which indicates how many clicks the dial should be rotated in that direction.

So, if the dial were pointing at 11, a rotation of R8 would cause the dial to point at 19. After that, a rotation of L19 would cause it to point at 0.

Because the dial is a circle, turning the dial left from 0 one click makes it point at 99. Similarly, turning the dial right from 99 one click makes it point at 0.

So, if the dial were pointing at 5, a rotation of L10 would cause it to point at 95. After that, a rotation of R5 could cause it to point at 0.

The dial starts by pointing at 50.

You could follow the instructions, but your recent required official North Pole secret entrance security training seminar taught you that the safe is actually a decoy. The actual password is the number of times the dial is left pointing at 0 after any rotation in the sequence.

For example, suppose the attached document contained the following rotations:

L68
L30
R48
L5
R60
L55
L1
L99
R14
L82
Following these rotations would cause the dial to move as follows:

The dial starts by pointing at 50.
The dial is rotated L68 to point at 82.
The dial is rotated L30 to point at 52.
The dial is rotated R48 to point at 0.
The dial is rotated L5 to point at 95.
The dial is rotated R60 to point at 55.
The dial is rotated L55 to point at 0.
The dial is rotated L1 to point at 99.
The dial is rotated L99 to point at 0.
The dial is rotated R14 to point at 14.
The dial is rotated L82 to point at 32.
Because the dial points at 0 a total of three times during this process, the password in this example is 3.

Analyze the rotations in your attached document. What's the actual password to open the door?
*/

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	startPosition := 50
	lines := readFile()
	//p1 := part1(startPosition, lines)
	p2 := part2(startPosition, lines)

	//fmt.Printf("The password in the first part is: %d\n", p1)
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

		rotations := 0
		oldPosition := currentPosition

		if distance > 99 {
			rotations = abs(distance) / 100
			fmt.Printf("the dial rotates %d since distance is %d, ", rotations, distance)
			distance = distance - (distance / 100 * 100)
			fmt.Printf("new distance: %d\n", distance)
		}

		password += rotations

		switch direction {
		case 'L':
			currentPosition = (currentPosition - distance + 100) % 100
		case 'R':
			currentPosition = (currentPosition + distance) % 100
		}

		switch {
		case direction == 'L' && currentPosition >= oldPosition:
			fmt.Printf("Line %s: Left rotation from %d to %d\n", line, oldPosition, currentPosition)
			password++
		case direction == 'R' && currentPosition <= oldPosition:
			fmt.Printf("Line %s: Right rotation from %d to %d\n", line, oldPosition, currentPosition)
			password++
		case currentPosition == 0:
			fmt.Printf("Line %s: Dial landed on 0 from %d\n", line, oldPosition)
			password++
		}
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

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
