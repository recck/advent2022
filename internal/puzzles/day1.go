package puzzles

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

type elf struct {
	calories []int
}

func (e *elf) addCalorie(calorie int) {
	e.calories = append(e.calories, calorie)
}

func (e *elf) totalCalories() int {
	calories := 0
	for i := range e.calories {
		calories += e.calories[i]
	}
	return calories
}

func newElf() elf {
	return elf{calories: []int{}}
}

type Day1 struct {
	elves []elf
}

func (d *Day1) addElf(e elf) {
	d.elves = append(d.elves, e)
}

func (d *Day1) Part1() any {
	maxCalories := -1

	for i := range d.elves {
		if d.elves[i].totalCalories() > maxCalories {
			maxCalories = d.elves[i].totalCalories()
		}
	}

	return maxCalories
}

func (d *Day1) Part2() any {
	elves := d.elves

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].totalCalories() > elves[j].totalCalories()
	})

	topSums := elves[0].totalCalories() + elves[1].totalCalories() + elves[2].totalCalories()
	return topSums
}

func (d *Day1) LoadInput(file *os.File) error {
	scanner := bufio.NewScanner(file)
	curElf := newElf()
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			d.addElf(curElf)
			curElf = newElf()
			continue
		}

		calories, _ := strconv.Atoi(line)
		curElf.addCalorie(calories)
	}
	d.addElf(curElf) // add the last elf

	return nil
}
