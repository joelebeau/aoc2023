package main

import (
    "fmt"
    "strings"
    "regexp"

    "github.com/joelebeau/aoc2023/aocUtils/input"
)

// My naming in today's problems is suboptimal, but I'm tired and busy today.
func main() {
    file, err := input.GetFile()
    if err != nil {
        panic(err)
    }
    defer file.Close()

    numMatcher := regexp.MustCompile(`\d+`)
    scanner, err := input.ReadInput(file)
    cardMatches := make([]int, 0)

    for scanner.Scan() {
        line := scanner.Text()
        lineSegments := strings.Split(line, "|")
        cardSegment := lineSegments[0]
        playerNumsSegment := lineSegments[1]
        cardNumStrings := strings.Split(cardSegment, ":")[1]

        cardNumbers := numMatcher.FindAllString(cardNumStrings, -1)
        playerNumbers := numMatcher.FindAllString(playerNumsSegment, -1)

        cardMatches = append(cardMatches, getMatchCount(cardNumbers, playerNumbers))
    }

    fmt.Println(getTotalCardsFromCardMatches(cardMatches))
}

// God help me for what I am doing here. Tired brain, GO.
func getTotalCardsFromCardMatches(cardMatches []int) int {
    // slice of arrays of int pairs. First element of each is card matches. Second is number of copies
    cardMatchesWithCounts := make([][2]int, len(cardMatches))
    for i, cardMatch := range cardMatches {
        cardMatchesWithCounts[i][0] = cardMatch
        cardMatchesWithCounts[i][1] = 1
    }

    // Loop over each original card
    for i := 0; i < len(cardMatchesWithCounts); i++ {
        // For each copy of the card at idx i...
        for j := 0; j < cardMatchesWithCounts[i][1]; j++ {
            kEnd := min(i + cardMatchesWithCounts[i][0] + 1, len(cardMatchesWithCounts))
            // Loop over the next original cards to add copies as needed,
            // starting at the next record and going til we add all our copies
            // or run out of cards to copy
            for k := i + 1; k < kEnd; k++ {
                cardMatchesWithCounts[k][1] += 1
            }
        }
    }
    total := 0
    for _, cardMatchesWithCount := range cardMatchesWithCounts {
        total += cardMatchesWithCount[1]
    }
    return total
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
