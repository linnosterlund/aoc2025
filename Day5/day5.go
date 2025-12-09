package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IDRange struct {
	StartID int
	EndID   int
}

func main() {
	ranges, available := getIDs()
	var freshAvailableIDs []int

	for _, a := range available {
		if checkID(a, ranges) {
			freshAvailableIDs = append(freshAvailableIDs, a)
		}
	}

	freshIDs := checkIDRanges(ranges)

	fmt.Printf("Part 1 - number of fresh ids: %d\n", len(freshAvailableIDs))
	fmt.Printf("Part 2 - number of fresh ids: %d\n", freshIDs)
}

func checkIDRanges(idRanges []IDRange) int {
	merged := mergeRanges(idRanges)

	total := 0
	for _, r := range merged {
		total += r.EndID - r.StartID + 1
	}
	return total
}

func mergeRanges(idRanges []IDRange) []IDRange {
	if len(idRanges) == 0 {
		return nil
	}

	sort.Slice(idRanges, func(i, j int) bool {
		return idRanges[i].StartID < idRanges[j].StartID
	})

	merged := []IDRange{idRanges[0]}

	for i := 1; i < len(idRanges); i++ {
		prev := &merged[len(merged)-1]
		current := idRanges[i]

		if current.StartID <= prev.EndID+1 {
			if current.EndID > prev.EndID {
				prev.EndID = current.EndID
			}
		} else {
			merged = append(merged, current)
		}
	}

	return merged
}

func checkID(id int, idRanges []IDRange) bool {
	for _, idRange := range idRanges {
		if id >= idRange.StartID && id <= idRange.EndID {
			return true
		}
	}
	return false
}

func getIDs() ([]IDRange, []int) {
	var iDRanges []IDRange
	var availableIDs []int

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("error reading input", err)
		return iDRanges, availableIDs
	}
	defer file.Close()

	s := bufio.NewScanner(file)
	isRanges := true
	for s.Scan() {
		row := s.Text()

		if strings.TrimSpace(row) == "" {
			isRanges = false
			continue
		}

		if isRanges {
			idNumbers := strings.Split(row, "-")

			if len(idNumbers) != 2 {
				fmt.Printf("invalid range format %s\n", row)
				continue
			}

			startID := getInt(strings.TrimSpace(idNumbers[0]))
			endID := getInt(strings.TrimSpace(idNumbers[1]))
			iDRanges = append(iDRanges, IDRange{StartID: startID, EndID: endID})
		} else {
			availableIDs = append(availableIDs, getInt(row))
		}

	}

	if err := s.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return iDRanges, availableIDs
	}

	return iDRanges, availableIDs
}

func getInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("error parsing string %s", s)
	}
	return n
}
