package main

/*
Identify invalid IDs from an input file
- ID format {nr}-{nr}
- separated by commas, first ID and last ID connected by dash

*/

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type IDRange struct {
	StartID    int
	EndID      int
	InvalidIDs []int
}

func main() {
	iDRanges := getRanges()
	sumInvalidIDs := 0

	for _, r := range iDRanges {
		for id := r.StartID; id <= r.EndID; id++ {
			if checkRepetitions(id) {
				r.InvalidIDs = append(r.InvalidIDs, id)
				sumInvalidIDs += id
			}
		}
	}
	fmt.Printf("Part 1: sum of invalid IDs is %d\n", sumInvalidIDs)

	sumInvalidIDs = 0
	for _, r := range iDRanges {
		for id := r.StartID; id <= r.EndID; id++ {
			if checkAllRepetitions(id) {
				r.InvalidIDs = append(r.InvalidIDs, id)
				sumInvalidIDs += id
			}
		}
	}
	fmt.Printf("Part 2: sum of invalid IDs is %d\n", sumInvalidIDs)
}

func checkRepetitions(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	if length%2 != 0 {
		return false
	}

	half := length / 2

	return s[:half] == s[half:]
}

func checkAllRepetitions(n int) bool {
	s := strconv.Itoa(n)
	length := len(s)

	for patternLen := 1; patternLen <= length/2; patternLen++ {
		if length%patternLen != 0 {
			continue
		}

		pattern := s[:patternLen]
		repetitions := length / patternLen
		repeated := strings.Repeat(pattern, repetitions)

		if repeated == s {
			return true
		}
	}
	return false
}

func getRanges() []IDRange {
	var iDRanges []IDRange

	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error reading input", err)
		return iDRanges
	}

	text := strings.TrimSpace(string(data))
	rangeParts := strings.SplitSeq(text, ",")

	for part := range rangeParts {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}

		idNumbers := strings.Split(part, "-")
		if len(idNumbers) != 2 {
			fmt.Printf("invalid range format %s\n", part)
			continue
		}

		startID, err1 := strconv.Atoi(strings.TrimSpace(idNumbers[0]))
		endID, err2 := strconv.Atoi(strings.TrimSpace(idNumbers[1]))
		if err1 != nil || err2 != nil {
			fmt.Printf("error parsing numbers in range %s\n", part)
			continue
		}

		iDRanges = append(iDRanges, IDRange{StartID: startID, EndID: endID, InvalidIDs: []int{}})
	}

	return iDRanges
}
