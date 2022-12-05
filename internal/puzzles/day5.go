package puzzles

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type stack struct {
	crates []string
}

func (s *stack) pushBack(item string) {
	s.crates = append(s.crates, item)
}

func (s *stack) pushFront(item string) {
	s.crates = append([]string{item}, s.crates...)
}

func (s *stack) pop() string {
	cL := len(s.crates)
	if cL == 0 {
		return ""
	}

	val := s.crates[cL-1]
	s.crates = s.crates[:cL-1]

	return val
}

func newStack() *stack {
	return &stack{crates: []string{}}
}

type instruction struct {
	quantity, from, to int
}

type Day5 struct {
	stacks       map[int]*stack
	instructions []instruction
}

func (d *Day5) getStackCopy() map[int]*stack {
	stackCopy := make(map[int]*stack)
	for i := range d.stacks {
		stackCopy[i] = newStack()
		for j := range d.stacks[i].crates {
			stackCopy[i].pushBack(d.stacks[i].crates[j])
		}
	}

	return stackCopy
}

func (d *Day5) Part1() any {
	s := d.getStackCopy()
	for i := range d.instructions {
		curInt := d.instructions[i]
		for j := 0; j < curInt.quantity; j++ {
			val := s[curInt.from].pop()
			if val == "" {
				continue
			}
			s[curInt.to].pushBack(val)
		}
	}

	var tops []string
	for m := 1; m <= len(s); m++ {
		tops = append(tops, s[m].pop())
	}

	return strings.Join(tops, "")
}

func (d *Day5) Part2() any {
	s := d.getStackCopy()
	for i := range d.instructions {
		curInt := d.instructions[i]
		var vals []string
		for j := 0; j < curInt.quantity; j++ {
			val := s[curInt.from].pop()
			if val == "" {
				continue
			}

			vals = append(vals, val)
		}
		for k := len(vals) - 1; k >= 0; k-- {
			s[curInt.to].pushBack(vals[k])
		}
	}

	var tops []string
	for m := 1; m <= len(s); m++ {
		tops = append(tops, s[m].pop())
	}

	return strings.Join(tops, "")
}

func (d *Day5) LoadInput(file *os.File) error {
	var instructions []instruction
	stacks := make(map[int]*stack)
	moveRegex := regexp.MustCompile(`^move (\d+) from (\d+) to (\d+)$`)
	crateRegex := regexp.MustCompile(`\[\w]`)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		}

		if moveRegex.MatchString(line) {
			matches := moveRegex.FindStringSubmatch(line)
			quantity, _ := strconv.Atoi(matches[1])
			from, _ := strconv.Atoi(matches[2])
			to, _ := strconv.Atoi(matches[3])
			instructions = append(instructions, instruction{
				quantity: quantity,
				from:     from,
				to:       to,
			})
		} else if crateRegex.MatchString(line) {
			sr := []rune(line)
			for i := 0; i < len(sr); i += 4 {
				chunk := string(sr[i : i+3])
				chunk = strings.ReplaceAll(chunk, "[", "")
				chunk = strings.ReplaceAll(chunk, "]", "")
				if len(strings.TrimSpace(chunk)) == 0 {
					continue
				}
				stackId := i/4 + 1
				// push front because the input is top to bottom
				if existingStack, ok := stacks[stackId]; ok {
					existingStack.pushFront(chunk)
				} else {
					stacks[stackId] = newStack()
					stacks[stackId].pushFront(chunk)
				}
			}
		}
	}

	d.instructions = instructions
	d.stacks = stacks

	return nil
}
