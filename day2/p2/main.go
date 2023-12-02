package main

import (
	"errors"
	"fmt"
	"strings"
	"strconv"
	"github.com/joelebeau/aoc2023/aocUtils/input"
)

type game struct {
	id int
	blues []int
	reds []int
	greens  []int
}

func main() {
	file, err := input.GetFile()
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner, err := input.ReadInput(file)
	if err != nil {
		panic(err)
	}

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		lineGame, err := buildGame(line)
		if err != nil {
			panic(err)
		}

		maxRed := max(lineGame.reds)
		maxGreen := max(lineGame.greens)
		maxBlue := max(lineGame.blues)

		power := maxRed * maxGreen * maxBlue

		total += power
	}

	fmt.Println(total)
}

func buildGame(line string) (*game, error) {
	var newGame game

	majorLineGroups := strings.Split(line, ":")
	gameIdStr := strings.Split(majorLineGroups[0], " ")
	gameId, err := strconv.Atoi(gameIdStr[1])
	if err != nil {
		return nil, err
	}
	newGame.id = gameId 
	
	// Get color groups
	rounds := strings.Split(majorLineGroups[1], ";")
	for _, round := range rounds {
		for _, colors := range strings.Split(round, ",") {
			colorSplit := strings.Split(colors, " ")
			// We get a space at the start of the line so this splits into 3 values
			numStr := colorSplit[1]
			num, err := strconv.Atoi(numStr)
			if err != nil {
				fmt.Println("Couldn't convert " + numStr + " to number")
				return nil, err
			}

			color := colorSplit[2]
			switch color {
			case "red":
				newGame.reds = append(newGame.reds, num)
			case "blue":
				newGame.blues = append(newGame.blues, num)
			case "green":
				newGame.greens = append(newGame.greens, num)
			default:
				return nil, errors.New(fmt.Sprintf("%v is an invalid color", color))
			}
		}
	}

	return &newGame, nil
}

func max(list []int) int {
	result := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] > result {
			result = list[i]
		}
	}

	return result
}
