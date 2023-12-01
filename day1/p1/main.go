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
	var firstAndLastNumericBytes [2]byte
	r := regexp.MustCompile(`\D`)

	strWithoutNonNumeric := r.ReplaceAllString(line, "")
	firstAndLastNumericBytes[0] = strWithoutNonNumeric[0]
	firstAndLastNumericBytes[1] = strWithoutNonNumeric[len(strWithoutNonNumeric)-1]

	return string(firstAndLastNumericBytes[:])
}
