package main

import (
    "fmt"

    "github.com/joelebeau/aoc2023/aocUtils/input"
)

func main() {
    file, err := input.GetFile()
    if err != nil {
        panic(err)
    }
    defer file.Close()

    scanner, err := input.ReadInput(file)
    
    for scanner.Scan() {
        line := scanner.Text()
        fmt.Println(line)
    }
}
