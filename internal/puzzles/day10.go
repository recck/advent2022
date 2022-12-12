package puzzles

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cycleInstruction int

const (
	cycleAdd cycleInstruction = iota
	cycleNoop
)

type day10instruction struct {
	instruction cycleInstruction
	value       int
}

type Day10 struct {
	instructions []day10instruction
}

func (_ *Day10) getValue(cycle, x int) int {
	if cycle > 220 {
		return 0
	}

	if cycle%20 == 0 && cycle%40 != 0 {
		return cycle * x
	}

	return 0
}

func (d *Day10) Part1() any {
	cycle := 0
	x := 1
	var curInstruction day10instruction
	seen := make(map[int]int)
	sums := 0

	for i := 0; i < len(d.instructions); {
		cycle++
		curInstruction = d.instructions[i]
		sums += d.getValue(cycle, x)
		if _, ok := seen[i]; !ok {
			seen[i] = 1
		} else {
			seen[i]++
		}

		if seen[i] == 2 && curInstruction.instruction == cycleAdd {
			x += curInstruction.value
			i++
		} else if curInstruction.instruction == cycleNoop {
			i++
		}
	}

	return sums
}

func (d *Day10) Part2() any {
	cycle := 0
	x := 1
	var curInstruction day10instruction
	crt := make([]string, 240)
	seen := make(map[int]int)

	for i := 0; i < len(d.instructions); {
		curInstruction = d.instructions[i]
		if _, ok := seen[i]; !ok {
			seen[i] = 1
		} else {
			seen[i]++
		}

		if x-1 <= cycle%40 && cycle%40 <= x+1 {
			crt[cycle] = "#"
		} else {
			crt[cycle] = "."
		}

		if seen[i] == 2 && curInstruction.instruction == cycleAdd {
			x += curInstruction.value
			i++
		} else if curInstruction.instruction == cycleNoop {
			i++
		}
		cycle++
	}

	output := ""
	for i := 0; i < len(crt); i++ {
		if i%40 == 0 {
			output += "\n"
		}
		output += fmt.Sprintf(" %s ", crt[i])
	}

	return output
}

func (d *Day10) LoadInput(file *os.File) error {
	var instructions []day10instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		v := 0
		op := cycleNoop
		chunks := strings.Fields(scanner.Text())
		if len(chunks) == 2 {
			v, _ = strconv.Atoi(chunks[1])
			op = cycleAdd
		}
		instructions = append(instructions, day10instruction{
			instruction: op,
			value:       v,
		})
	}

	d.instructions = instructions
	return nil
}
