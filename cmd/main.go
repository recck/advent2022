package main

import (
	"fmt"
	"github.com/recck/advent2022/internal/puzzles"
	"os"
	"strconv"
)

var puzzleList = map[int]puzzles.Puzzle{
	1: &puzzles.Day1{},
	2: &puzzles.Day2{},
	3: &puzzles.Day3{},
	4: &puzzles.Day4{},
	5: &puzzles.Day5{},
	6: &puzzles.Day6{},
}

func main() {
	inputArgs := os.Args[1:]

	if len(inputArgs) != 2 {
		fmt.Println("missing args")
		os.Exit(1)
	}

	day := inputArgs[0]
	dayInt, err := strconv.Atoi(day)

	if err != nil {
		fmt.Printf("day should be a number, %v\n", err)
		os.Exit(1)
	}

	file, err := os.Open(inputArgs[1])

	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}

	curPuzzle, ok := puzzleList[dayInt]

	if !ok {
		fmt.Println("puzzle does not exist")
		os.Exit(1)
	}

	puzzles.Solve(curPuzzle, file)
}
