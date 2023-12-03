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

		nums, ranges, err := findNumbersWithIndexes(currLine)
		if err != nil {
			panic(err)
		}

		total += getTotalFromHits(currLine, prevLine, nextLine, nums, ranges)
	}
	// final iteration for the last line
	prevLine = currLine
	currLine = nextLine
	nextLine = ""

	nums, ranges, err := findNumbersWithIndexes(currLine)
	if err != nil {
		panic(err)
	}

	total += getTotalFromHits(currLine, prevLine, nextLine, nums, ranges)

	fmt.Println("Total is", total)
}

// Returns the number string found and
func findNumbersWithIndexes(line string) ([]int, [][]int, error) {
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

func getTotalFromHits(currLine string, prevLine string, nextLine string, nums []int, ranges [][]int) int {
	subTotal := 0
	for idx, num := range nums {
		isHit := checkCurrLine(currLine, ranges[idx])
		if !isHit {
			isHit = isHit || checkAdjacentLine(prevLine, ranges[idx])
		}

		if !isHit {
			isHit = isHit || checkAdjacentLine(nextLine, ranges[idx])
		}

		if isHit {
			fmt.Println(num)
			subTotal += num
		}
	}

	return subTotal
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
