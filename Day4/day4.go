package main

/*
Position of paper roll:
*/

import (
	"bufio"
	"fmt"
	"os"
)

type PositionState int

const (
	empty PositionState = iota
	roll
)

type Position struct {
	X int
	Y int
}

var dir = []Position{
	{-1, -1},
	{-1, 0},
	{-1, 1},
	{0, -1},
	{0, 1},
	{1, -1},
	{1, 0},
	{1, 1},
}

func main() {
	positions := getFile()

	accessible := part1(positions)
	totalAccessible := part2(positions)

	fmt.Printf("Part 1: %d rolls of papers can be removed\n", accessible)
	fmt.Printf("Part 2: %d rolls of papers can be removed in total\n", totalAccessible)
}

func part2(ps map[Position]PositionState) int {
	var accessibleRolls []Position
	totalAccessible := 0
	accessible := 0

	for {
		for p, value := range ps {
			if value == 1 {
				neighbourRolls := checkNeighbours(p, ps)
				if neighbourRolls < 4 {
					accessible++
					accessibleRolls = append(accessibleRolls, p)
				}
			}
		}

		if accessible == 0 {
			break
		}

		totalAccessible += accessible
		accessible = 0
		for _, pos := range accessibleRolls {
			ps[pos] = empty
		}
	}

	return totalAccessible
}

func part1(ps map[Position]PositionState) int {
	accessible := 0

	for p, value := range ps {
		if value == 1 {
			neighbourRolls := checkNeighbours(p, ps)
			if neighbourRolls < 4 {
				accessible++
			}
		}
	}

	return accessible
}

func checkNeighbours(p Position, ps map[Position]PositionState) int {
	neighbourRolls := 0
	for _, d := range dir {
		neighbour := Position{
			X: p.X + d.X,
			Y: p.Y + d.Y,
		}

		neighbourVal, exists := ps[neighbour]
		if exists && neighbourVal == 1 {
			neighbourRolls++
		}
	}
	return neighbourRolls
}

func getFile() map[Position]PositionState {
	positions := make(map[Position]PositionState)

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error reading input", err)
		return positions
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	y := 0

	for s.Scan() {
		row := s.Text()
		for x, ch := range row {
			pos := Position{X: x, Y: y}

			switch ch {
			case '@':
				positions[pos] = roll
			case '.':
				positions[pos] = empty
			}
		}
		y++
	}

	if err := s.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return positions
	}

	return positions
}
