package puzzles

import (
	"fmt"
	"os"
)

type Puzzle interface {
	Part1() any
	Part2() any
	LoadInput(file *os.File) error
}

func Solve(p Puzzle, file *os.File) {
	if err := p.LoadInput(file); err != nil {
		fmt.Printf("error loading input: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("part 1: %v\npart 2: %v\n", p.Part1(), p.Part2())
}
