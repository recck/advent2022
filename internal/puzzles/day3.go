package puzzles

import (
	"bufio"
	"os"
	"strings"
	"unicode"
)

func getPriority(r rune) int {
	if unicode.IsUpper(r) {
		return int(r) - 38
	}
	return int(r) - 96
}

func getOverlappingLetter(stringList ...string) rune {
	s1 := stringList[0]
nextChar:
	for c := range s1 {
		for i := 1; i < len(stringList); i++ {
			if !strings.Contains(stringList[i], string(s1[c])) {
				continue nextChar
			}
		}
		return rune(s1[c])
	}

	return 0
}

type rucksack struct {
	firstCompartment, secondCompartment, both string
}

type Day3 struct {
	rucksacks []rucksack
}

func (d *Day3) Part1() any {
	sum := 0
	for r := range d.rucksacks {
		letter := getOverlappingLetter(d.rucksacks[r].firstCompartment, d.rucksacks[r].secondCompartment)
		sum += getPriority(letter)
	}

	return sum
}

func (d *Day3) Part2() any {
	sum := 0
	for i := 0; i < len(d.rucksacks); i += 3 {
		letter := getOverlappingLetter(d.rucksacks[i].both, d.rucksacks[i+1].both, d.rucksacks[i+2].both)
		sum += getPriority(letter)
	}

	return sum
}

func (d *Day3) LoadInput(file *os.File) error {
	var rucksacks []rucksack
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cSize := len(line) / 2
		rucksacks = append(rucksacks, rucksack{
			firstCompartment:  line[:cSize],
			secondCompartment: line[cSize:],
			both:              line,
		})
	}
	d.rucksacks = rucksacks

	return nil
}
