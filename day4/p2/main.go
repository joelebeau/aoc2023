package main

import (
    "fmt"
    "strings"
    "regexp"

    "github.com/joelebeau/aoc2023/aocUtils/input"
)

func main() {
    file, err := input.GetFile()
    if err != nil {
        panic(err)
    }
    defer file.Close()

    numMatcher := regexp.MustCompile(`\d+`)
    scanner, err := input.ReadInput(file)

    total := 0
    for scanner.Scan() {
        line := scanner.Text()
        lineSegments := strings.Split(line, "|")
        cardSegment := lineSegments[0]
        playerNumsSegment := lineSegments[1]
        cardNumStrings := strings.Split(cardSegment, ":")[1]

        cardNumbers := numMatcher.FindAllString(cardNumStrings, -1)
        playerNumbers := numMatcher.FindAllString(playerNumsSegment, -1)


        total += cardTotal
    }

    fmt.Println(total)
}

func initOrDoubleScore(currScore int) int {
    if currScore == 0 {
        return 1
    }

    return currScore * 2
}

func getMatchCount(cardNumbers []string, playerNumbers []string) int {
    var count int
    cardMap := make(map[string]bool)
    for _, num := range cardNumbers {
        cardMap[num] = true
    }

    for _, num := range playerNumbers {
        if cardMap[num] {
            count += 1
        }
    }

    return count
}
