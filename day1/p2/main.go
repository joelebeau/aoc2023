package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	file, err := getFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner, err := readInput(file)
	if err != nil {
		panic(err)
	}

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		numStr := findNumberPair(line)
		num, err := strconv.Atoi(numStr)
		if err != nil {
			// We can't finish the calculation, just panic
			panic(err)
		}

		total += num
	}

	fmt.Println(total)
}

func getFile() (*os.File, error) {
	filePath := os.Args[1]

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func readInput(file *os.File) (*bufio.Scanner, error) {
	fileInfo, err := file.Stat()
	if err != nil {
		return nil, err
	}

	fileSize := fileInfo.Size()
	scanner := bufio.NewScanner(file)

	buffer := make([]byte, fileSize)
	scanner.Buffer(buffer, int(fileSize))

	return bufio.NewScanner(file), nil
}

func findNumberPair(line string) string {
	// Mapping to strings only to limit the rework from part one.
	strToNumStrMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"zero":  "0",
		"1":     "1",
		"2":     "2",
		"3":     "3",
		"4":     "4",
		"5":     "5",
		"6":     "6",
		"7":     "7",
		"8":     "8",
		"9":     "9",
		"0":     "0",
	}

	var firstAndLastNumStr [2]string
	r := regexp.MustCompile(`([0-9]|one|two|three|four|five|six|seven|eight|nine|zero)`)
	matches := r.FindAllString(line, -1)

	firstAndLastNumStr[0] = strToNumStrMap[matches[0]]
	firstAndLastNumStr[1] = strToNumStrMap[matches[len(matches)-1]]

	return firstAndLastNumStr[0] + firstAndLastNumStr[1]
}
