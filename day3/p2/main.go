package main

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/joelebeau/aoc2023/aocUtils/input"
)

func main() {
	file, err := input.GetFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner, err := input.ReadInput(file)

	var prevLine, currLine, nextLine string

	total := 0
	// prep the first line
	scanner.Scan()
	nextLine = "INIT" // Just a hack to start the loop
	for nextLine != "" {
		prevLine = currLine
		currLine = scanner.Text()
		if scanner.Scan() {
			nextLine = scanner.Text()
		} else {
			nextLine = ""
		}

		astIndexes := findAsteriskIndexes(currLine)

		iterTotal, err := getTotalFromHits(currLine, prevLine, nextLine, astIndexes)
		if err != nil {
			panic(err)
		}
		total += iterTotal
	}
	// final iteration for the last line
	prevLine = currLine
	currLine = nextLine
	nextLine = ""

	astIndexes := findAsteriskIndexes(currLine)
	iterTotal, err := getTotalFromHits(currLine, prevLine, nextLine, astIndexes)
	if err != nil {
		panic(err)
	}

	total += iterTotal

	fmt.Println("Total is", total)
}

func findAsteriskIndexes(line string) []int {
	asteriskMatcher := regexp.MustCompile(`\*`)
	matches := asteriskMatcher.FindAllStringIndex(line, -1)
	simplifiedMatches := make([]int, len(matches))

	for idx, m := range matches {
		simplifiedMatches[idx] = m[0]
	}

	return simplifiedMatches
}

// Returns the number string found and the locations in the string where they matched
func findNumbersWithIndexes(line string) ([]int, [][]int, error) {
	if line == "" {
		return []int{}, [][]int{}, nil
	}
	numMatcher := regexp.MustCompile(`\d+`)
	matches := numMatcher.FindAllStringIndex(line, -1)
	nums := make([]int, len(matches))
	for idx, matchRange := range matches {
		numStr := line[matchRange[0]:matchRange[1]]
		num, err := strconv.Atoi(numStr)
		if err != nil {
			return nil, nil, err
		}
		nums[idx] = num
	}

	return nums, matches, nil
}

func getTotalFromHits(currLine string, prevLine string, nextLine string, astIndexes []int) (int, error) {
	plNums, plRanges, err := findNumbersWithIndexes(prevLine)
	if err != nil {
		return -1, err
	}
	clNums, clRanges, err := findNumbersWithIndexes(currLine)
	if err != nil {
		return -1, err
	}
	nlNums, nlRanges, err := findNumbersWithIndexes(nextLine)
	if err != nil {
		return -1, err
	}
	subTotal := 0
	for _, astIdx := range astIndexes {
		minIdx := max(0, astIdx-1)
		maxIdx := min(len(currLine), astIdx+1)

		foundNums := make([]int, 0)

		for plIdx, plRange := range plRanges {
			if doRangesOverlap([]int{minIdx, maxIdx}, plRange) {
				foundNums = append(foundNums, plNums[plIdx])
			}
		}
		for clIdx, clRange := range clRanges {
			if doRangesOverlap([]int{minIdx, maxIdx}, clRange) {
				foundNums = append(foundNums, clNums[clIdx])
			}
		}
		for nlIdx, nlRange := range nlRanges {
			if doRangesOverlap([]int{minIdx, maxIdx}, nlRange) {
				foundNums = append(foundNums, nlNums[nlIdx])
			}
		}

		if len(foundNums) > 2 {
			panic("Found too many matches")
		}
		if len(foundNums) == 2 {
			subTotal += (foundNums[0] * foundNums[1])
		}
	}

	return subTotal, nil
}

func doRangesOverlap(r1 []int, r2 []int) bool {
	// r2[1] is exclusive, so we need to subtract 1
	return r1[0] <= r2[1]-1 && r2[0] <= r1[1]
}

func checkAdjacentLine(adjLine string, matchRange []int) bool {
	symbolMatcher := regexp.MustCompile(`[^0-9\.]`)
	if adjLine == "" {
		return false
	}

	startIdx := max(0, matchRange[0]-1)
	endIdx := min(len(adjLine), matchRange[1]+1)

	adjLineSubStr := adjLine[startIdx:endIdx]
	matches := symbolMatcher.MatchString(adjLineSubStr)

	return matches
}

func checkCurrLine(line string, matchRange []int) bool {
	symbolMatcher := regexp.MustCompile(`[^0-9\.]`)
	bytesToCheck := make([]byte, 0)

	// Add the location to the left if it's not the first item in the line
	if matchRange[0] > 0 {
		bytesToCheck = append(bytesToCheck, line[matchRange[0]-1])
	}

	// Add the location to the right if it's not the last item in the line
	if matchRange[1] < len(line)-2 {
		bytesToCheck = append(bytesToCheck, line[matchRange[1]])
	}

	matches := symbolMatcher.Match(bytesToCheck)

	return matches
}
